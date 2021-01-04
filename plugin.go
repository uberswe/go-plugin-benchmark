package benchmark

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

// RandIntResponder is a responder for hashicorp/go-plugin
type RandIntResponder interface {
	Respond() int
}

// RandIntRPC is a struct used for hashicorp/go-plugin
type RandIntRPC struct{ client *rpc.Client }

// Respond is a function used while benchmarking hashicorp/go-plugin
func (g *RandIntRPC) Respond() int {
	var resp int
	err := g.client.Call("Plugin.Respond", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

// RandIntRPCServer is the RPC server used when benchmarking hashicorp/go-plugin
type RandIntRPCServer struct {
	Impl RandIntResponder
}

// Respond is a function used while benchmarking hashicorp/go-plugin
func (s *RandIntRPCServer) Respond(args interface{}, resp *int) error {
	*resp = s.Impl.Respond()
	return nil
}

// RandIntPlugin is a struct used while benchmarking hashicorp/go-plugin
type RandIntPlugin struct {
	Impl RandIntResponder
}

// Server is the server used while benchmarking hashicorp/go-plugin
func (p *RandIntPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RandIntRPCServer{Impl: p.Impl}, nil
}

// Client is the client used while benchmarking hashicorp/go-plugin
func (RandIntPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RandIntRPC{client: c}, nil
}

type plug struct {
	client *rpc.Client
}

// RandInt is a function used for testing natefinch/pie
func (p plug) RandInt(in int) (result int, err error) {
	err = p.client.Call("Plugin.RandInt", in, &result)
	return result, err
}
