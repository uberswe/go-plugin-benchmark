# Golang Plugin Benchmark

A comparison of the [go plugin package](https://golang.org/pkg/plugin/) and other plugin implementations for golang. The reason for this is to evaluate plugin options for [Beubo](https://github.com/uberswe/beubo) which is a CMS written in go. As [shown in a previous benchmark](https://github.com/uberswe/goplugins) there is basically no difference in performance of using go plugins to running a function directly in code.

The following packages have been benchmarked.

 - [go plugin package](https://golang.org/pkg/plugin/)
 - [hashicorp/go-plugin](https://github.com/hashicorp/go-plugin)
 - [natefinch/pie](https://github.com/natefinch/pie)
 - [dullgiulio/pingo](https://github.com/dullgiulio/pingo)
 - [elliotmr/plug](https://github.com/elliotmr/plug)
 - [traefik/yaegi](https://github.com/traefik/yaegi)
 - [pkujhd/goloader](https://github.com/pkujhd/goloader)
 
Several of the other packages use RPC or similar methods instead of the go plugin package which gets around issues such as, but not limited to, [not being compatible with Windows](https://github.com/golang/go/issues/19282) and [package paths and GOPATH needing to be the same between apps and plugins](https://github.com/golang/go/issues/20481).

With the addition of [Yaegi](https://github.com/traefik/yaegi), I am also benchmarking interpreters. 

**Do you know any other plugin packages?** Please open an [issue](https://github.com/uberswe/go-plugin-benchmark/issues/new) or pull request.

**Found an issue with a benchmark?** Please open an [issue](https://github.com/uberswe/go-plugin-benchmark/issues/new) or pull request.

## Setup

The Goloader package requires the following to work
```
cp -r $GOROOT/src/cmd/internal $GOROOT/src/cmd/objfile
```

## Results

Most plugins tested are using RPC which adds about 30 - 50 microseconds to plugin calls (or 0.03 - 0.05 milliseconds) over the golang plugin package.

The [goloader](https://github.com/pkujhd/goloader) package is interesting and may provide a good alternative to the go plugin package. One drawback is that it uses internal packages which requires renaming the internal folder locally and I have not tested compatibility to see if it solves the problems with the go plugin package.

## Benchmarks

| Name                       | Operations |       ns/op |
| -------------------------- |:----------:|------------:|
| go plugin package          |  87055964  | 13.61 ns/op |
| hashicorp/go-plugin        |   26644    | 46741 ns/op |
| natefinch/pie              |   29536    | 37524 ns/op |
| dullgiulio/pingo over tcp  |   19928    | 55858 ns/op |
| dullgiulio/pingo over unix |   30532    | 39308 ns/op |
| elliotmr/plug              |   90650    | 12559 ns/op |
| traefik/yaegi              |  1478181   | 822.1 ns/op |
| pkujhd/goloader            |  86624695  | 14.15 ns/op |

Last run on May 28th, 2022 with Go version 1.18. Benchmark performed on a MacBook Pro (15-inch, 2018) with a 2,9 GHz 6-Core Intel Core i9 processor and 32 GB 2400 MHz DDR4 ram.

```
% bash run.sh
goos: darwin
goarch: amd64
pkg: github.com/uberswe/go-plugin-benchmark
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkPluginRandInt/golang-plugin-12         87055964                13.61 ns/op
BenchmarkHashicorpGoPluginRandInt/hashicorp-go-plugin-12                   26644             46741 ns/op
BenchmarkPieRandInt/pie-12                                                 29536             37524 ns/op
BenchmarkPingoTcpRandInt/pingo-tcp-12                                      19928             55858 ns/op
BenchmarkPingoTcpRandInt/pingo-unix-12                                     30532             39308 ns/op
BenchmarkPlugRandInt/plug-12                                               90650             12559 ns/op
BenchmarkYaegiRandInt/yaegi-12                                           1478181               822.1 ns/op
BenchmarkGoloaderRandInt/goloader-12                                    86624695                14.15 ns/op
PASS
ok      github.com/uberswe/go-plugin-benchmark  14.301s
```

## Contributing

If you have a plugin or interpreter written in go which you would like to benchmark feel free to [open an issue](https://github.com/uberswe/go-plugin-benchmark/issues/new).

Would you like to add more benchmarks? Please feel free to fork this repository and open a pull request with the updated changes. Please make sure to add any needed code and also update the Readme. Your code belongs to you but it must fall under the same LICENSE as this repository to be included.