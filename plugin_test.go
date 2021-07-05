package benchmark

import (
	"github.com/dullgiulio/pingo"
	"github.com/hashicorp/go-hclog"
	hashicorpplugin "github.com/hashicorp/go-plugin"
	"github.com/natefinch/pie"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	plugplugin "github.com/uberswe/go-plugin-benchmark/plug"
	"net/rpc/jsonrpc"
	"os"
	"os/exec"
	"plugin"
	"runtime"
	"testing"
)

// BenchmarkPluginRandInt uses a go plugin and tests math/rand for generating random integers
func BenchmarkPluginRandInt(b *testing.B) {
	plug, err := plugin.Open("./plugin.so")
	if err != nil {
		panic(err)
	}

	randInt, err := plug.Lookup("RandInt")
	if err != nil {
		panic(err)
	}

	randFunc, ok := randInt.(func() int)
	if !ok {
		panic("unexpected type from module symbol")
	}

	b.Run("golang-plugin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			randFunc()
		}
	})
}

func BenchmarkHashicorpGoPluginRandInt(b *testing.B) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Off,
	})

	var handshakeConfig = hashicorpplugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "RAND_PLUGIN",
		MagicCookieValue: "int",
	}

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]hashicorpplugin.Plugin{
		"responder": &RandIntPlugin{},
	}

	// We're a host! Start by launching the plugin process.
	client := hashicorpplugin.NewClient(&hashicorpplugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./hashicorpgoplugin"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		panic(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("responder")
	if err != nil {
		panic(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	intResponder := raw.(RandIntResponder)
	b.Run("hashicorp-go-plugin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			intResponder.Respond()
		}
	})
}

func BenchmarkPieRandInt(b *testing.B) {
	path := "./pieplugin"
	if runtime.GOOS == "windows" {
		path = path + ".exe"
	}

	client, err := pie.StartProviderCodec(jsonrpc.NewClientCodec, os.Stderr, path)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	p := plug{client}
	b.Run("pie", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = p.RandInt(i)
		}
	})
}

func BenchmarkPingoTcpRandInt(b *testing.B) {
	p := pingo.NewPlugin("tcp", "./pingoplugin")
	p.Start()

	var resp int

	b.Run("pingo-tcp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p.Call("MyPlugin.RandInt", 1, &resp)
		}
	})
	p.Stop()

	p2 := pingo.NewPlugin("unix", "./pingoplugin")
	p2.SetSocketDirectory("./")

	p2.Start()

	b.Run("pingo-unix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p2.Call("MyPlugin.RandInt", 1, &resp)
		}
	})
	p2.Stop()
}

func BenchmarkPlugRandInt(b *testing.B) {
	service, _ := plugplugin.LoadRandomIntService("./plugplugin")
	b.Run("plug", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = service.Get(1)
		}
	})
}

func BenchmarkYaegiRandInt(b *testing.B) {
	var src = `package test
import "math/rand"
func RandInt(i int) int { return rand.Int() }`

	i := interp.New(interp.Options{})

	// To handle import of "math/rand"
	i.Use(stdlib.Symbols)

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("test.RandInt")
	if err != nil {
		panic(err)
	}

	randIntFunc := v.Interface().(func(int) int)
	b.Run("yaegi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = randIntFunc(1)
		}
	})
}
