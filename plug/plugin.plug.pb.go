// Code generated by protoc-gen-plug. DO NOT EDIT.

package plug

import (
	fmt "fmt"
	runtime "github.com/elliotmr/plug/pkg/runtime"
	proto "google.golang.org/protobuf/proto"
)

const (
	RandomIntServiceGetService       runtime.Service = 0x0000
	RandomIntServiceHandshakeService runtime.Service = 0x00FF
)

type RandomIntService interface {
	Get(key int64) (int64, error)
}

type randomintservicePlugin struct {
	Impl RandomIntService
}

func RunRandomIntService(impl RandomIntService) error {
	s := &randomintservicePlugin{Impl: impl}
	return runtime.Run(s, "magic", 1, 0)
}

func (x *randomintservicePlugin) Link(srv runtime.Service) (proto.Message, runtime.GenPluginMethod, error) {
	switch srv {
	case RandomIntServiceGetService:
		return &RandIntRequest{}, x.Get, nil
	}
	return nil, nil, fmt.Errorf("unknown service: %d", srv)
}

func (x *randomintservicePlugin) Get(req proto.Message) (proto.Message, error) {
	in, ok := req.(*RandIntRequest)
	if !ok {
		return nil, fmt.Errorf("invalid request type")
	}
	value, err := x.Impl.Get(in.Key)
	return &RandIntResponse{Value: value}, err
}

type randomintserviceHost struct {
	c *runtime.Host
}

func LoadRandomIntService(s string) (RandomIntService, error) {
	c, err := runtime.Load(s, "magic", 1, 0)
	if err != nil {
		return nil, fmt.Errorf("unable to load plugin: %w", err)
	}
	return &randomintserviceHost{c: c}, nil
}

func (x *randomintserviceHost) Get(key int64) (int64, error) {
	resp := &RandIntResponse{}
	err := x.c.SendRecv(RandomIntServiceGetService, &RandIntRequest{Key: key}, resp)
	return resp.Value, err
}

func TestRandomIntService(impl RandomIntService) (RandomIntService, error) {
	s := &randomintservicePlugin{Impl: impl}
	c, err := runtime.Test(s, "magic", 1, 0)
	if err != nil {
		return nil, fmt.Errorf("unable to load plugin: %w", err)
	}
	return &randomintserviceHost{c: c}, nil
}
