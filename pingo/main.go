package main

import (
	"github.com/dullgiulio/pingo"
	"math/rand"
)

// MyPlugin is the plugin struct
type MyPlugin struct{}

// RandInt returns a random int
func (p *MyPlugin) RandInt(in int, msg *int) error {
	*msg = rand.Int()
	return nil
}

func main() {
	plugin := &MyPlugin{}
	pingo.Register(plugin)
	_ = pingo.Run()
}
