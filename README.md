# Golang Plugin Benchmark

A comparison of the [go plugin package](https://golang.org/pkg/plugin/) and other plugin implementations for golang. The reason for this is to evaluate plugin options for [Beubo](https://github.com/uberswe/beubo) which is a CMS written in go. As [shown in a previous benchmark](https://github.com/uberswe/goplugins) there is basically no difference in performance of using go plugins to running a function directly in code.

The following packages have been benchmarked.

 - [go plugin package](https://golang.org/pkg/plugin/)
 - [hashicorp/go-plugin](https://github.com/hashicorp/go-plugin)
 - [natefinch/pie](https://github.com/natefinch/pie)
 - [dullgiulio/pingo](https://github.com/dullgiulio/pingo)
 - [elliotmr/plug](https://github.com/elliotmr/plug)
 - [traefik/yaegi](https://github.com/traefik/yaegi)
 
Several of the other packages use RPC or similar methods instead of the go plugin package which gets around issues such as, but not limited to, [not being compatible with Windows](https://github.com/golang/go/issues/19282) and [package paths and GOPATH needing to be the same between apps and plugins](https://github.com/golang/go/issues/20481).

With the addition of [Yaegi](https://github.com/traefik/yaegi), I am also benchmarking interpreters. 

**Do you know any other plugin packages?** Please open an [issue](https://github.com/uberswe/go-plugin-benchmark/issues/new) or pull request.

**Found an issue with a benchmark?** Please open an [issue](https://github.com/uberswe/go-plugin-benchmark/issues/new) or pull request.

## Results

Most plugins tested are using RPC which adds about 30 - 50 microseconds to plugin calls (or 0.03 - 0.05 milliseconds) over the golang plugin package.

## Benchmarks

| Name                       | Operations   | ns/op       |
| -------------------------- |:------------:| -----------:|
| go plugin package          | 76898038     | 13.89 ns/op |
| hashicorp/go-plugin        | 24984        | 42563 ns/op |
| natefinch/pie              | 33962        | 37022 ns/op |
| dullgiulio/pingo over tcp  | 22694        | 53541 ns/op |
| dullgiulio/pingo over unix | 29288        | 36140 ns/op |
| elliotmr/plug              | 96416        | 12674 ns/op |
| traefik/yaegi              | 1771837      | 713.4 ns/op |

Last run on July 5th, 2021. Benchmark performed on a MacBook Pro (15-inch, 2018) with a 2,9 GHz 6-Core Intel Core i9 processor and 32 GB 2400 MHz DDR4 ram.

```
% bash run.sh 
goos: darwin
goarch: amd64
pkg: github.com/uberswe/go-plugin-benchmark
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkPluginRandInt/golang-plugin-12         76898038                13.89 ns/op
BenchmarkHashicorpGoPluginRandInt/hashicorp-go-plugin-12                   24984             42563 ns/op
BenchmarkPieRandInt/pie-12                                                 33962             37022 ns/op
BenchmarkPingoTcpRandInt/pingo-tcp-12                                      22694             53541 ns/op
BenchmarkPingoTcpRandInt/pingo-unix-12                                     29288             36140 ns/op
BenchmarkPlugRandInt/plug-12                                               96416             12674 ns/op
BenchmarkYaegiRandInt/yaegi-12                                           1771837               713.4 ns/op
PASS
ok      github.com/uberswe/go-plugin-benchmark  11.760s
```

## Contributing

If you have a plugin or interpreter written in go which you would like to benchmark feel free to [open an issue](https://github.com/uberswe/go-plugin-benchmark/issues/new).

Would you like to add more benchmarks? Please feel free to fork this repository and open a pull request with the updated changes. Please make sure to add any needed code and also update the Readme. Your code belongs to you but it must fall under the same LICENSE as this repository to be included.