package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	benchmark "github.com/uberswe/go-plugin-benchmark"
	"math/rand"
	"os"
)

// RandIntResponder is a plugin struct
type RandIntResponder struct{}

// Respond is a plugin function returning a random int
func (g *RandIntResponder) Respond() int {
	return rand.Int()
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "RAND_PLUGIN",
	MagicCookieValue: "int",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Off,
	})

	responder := &RandIntResponder{}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"responder": &benchmark.RandIntPlugin{Impl: responder},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Logger:          logger,
	})
}
