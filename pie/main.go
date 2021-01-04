package main

import (
	"github.com/natefinch/pie"
	"math/rand"
	"net/rpc/jsonrpc"
)

func main() {
	p := pie.NewProvider()
	if err := p.RegisterName("Plugin", api{}); err != nil {
		panic(err)
	}
	p.ServeCodec(jsonrpc.NewServerCodec)
}

type api struct{}

func (api) RandInt(in int, response *int) error {
	*response = rand.Int()
	return nil
}
