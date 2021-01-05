# Golang Plugin Benchmark

A comparison of the [go plugin package](https://golang.org/pkg/plugin/) and other plugin implementations for golang. The reason for this is to evaluate plugin options for [Beubo](https://github.com/uberswe/beubo) which is a CMS written in go. As [shown in a previous benchmark](https://github.com/uberswe/goplugins) there is basically no difference in performance of using go plugins to running a function directly in code.

The following packages have been benchmarked.

 - [go plugin package](https://golang.org/pkg/plugin/)
 - [hashicorp/go-plugin](https://github.com/hashicorp/go-plugin)
 - [natefinch/pie](https://github.com/natefinch/pie)
 - [dullgiulio/pingo](https://github.com/dullgiulio/pingo)
 - [elliotmr/plug](https://github.com/elliotmr/plug)
 
Several of the other packages use RPC or similar methods instead of the go plugin package which gets around issues such as, but not limited to, [not being compatible with Windows](https://github.com/golang/go/issues/19282) and [package paths and GOPATH needing to be the same between apps and plugins](https://github.com/golang/go/issues/20481).

**Do you know any other plugin packages?** Please open an [issue](https://github.com/uberswe/go-plugin-benchmark/issues/new) or pull request.

**Found an issue with a benchmark?** Please open an [issue](https://github.com/uberswe/go-plugin-benchmark/issues/new) or pull request.

## Results

Most plugins tested are using RPC which adds about 30 - 50 microseconds to plugin calls (or 0.03 - 0.05 milliseconds) over the golang plugin package.

## Benchmarks

| Name                       | Operations  | ns/op       |
| -------------------------- |:-----------:| -----------:|
| go plugin package          | 85481671    | 13.8 ns/op  |
| hashicorp/go-plugin        | 27888       | 43508 ns/op |
| natefinch/pie              | 33175       | 36892 ns/op |
| dullgiulio/pingo over tcp  | 19911       | 60874 ns/op |
| dullgiulio/pingo over unix | 31688       | 37106 ns/op |
| elliotmr/plug              | 101408      | 12022 ns/op |

Benchmark was performed on a MacBook Pro (15-inch, 2018) with a 2,9 GHz 6-Core Intel Core i9 processor and 32 GB 2400 MHz DDR4 ram.

```
% bash run.sh
goos: darwin
goarch: amd64
pkg: github.com/uberswe/go-plugin-benchmark
BenchmarkPluginRandInt/golang-plugin-12         85481671                13.8 ns/op
BenchmarkHashicorpGoPluginRandInt/hashicorp-go-plugin-12                   27888             43508 ns/op
BenchmarkPieRandInt/pie-12                                                 33175             36892 ns/op
BenchmarkPingoTcpRandInt/pingo-tcp-12                                      19911             60874 ns/op
BenchmarkPingoTcpRandInt/pingo-unix-12                                     31688             37106 ns/op
BenchmarkPlugRandInt/plug-12                                              101408             12022 ns/op
PASS
ok      github.com/uberswe/go-plugin-benchmark  9.947s
```
