# Golang Plugin Benchmark

A comparison of the [go plugin package](https://golang.org/pkg/plugin/) and other plugin implementations for golang. The reason for this is to compare alternatives from a performance perspective. As [shown in a previous benchmark](https://github.com/uberswe/goplugins) there is basically no difference in performance of using go plugins to running a function directly in code.

## Benchmarks

| Name                                                          | Operations (higher is better) | ns/op (lower is better) |    type     |
|---------------------------------------------------------------|:-----------------------------:|------------------------:|:-----------:|
| [go plugin package](https://golang.org/pkg/plugin/)           |           44219324            |             30.35 ns/op |   native    |
| [hashicorp/go-plugin](https://github.com/hashicorp/go-plugin) |             3682              |            413257 ns/op |     rpc     |
| [natefinch/pie](https://github.com/natefinch/pie)             |             3933              |            328025 ns/op |     rpc     |
| [dullgiulio/pingo](https://github.com/dullgiulio/pingo)       |             4197              |            329354 ns/op |     tcp     |
| [dullgiulio/pingo](https://github.com/dullgiulio/pingo)       |             3110              |            465628 ns/op |    unix     |
| [elliotmr/plug](https://github.com/elliotmr/plug)             |             7998              |            162677 ns/op |     ipc     |
| [traefik/yaegi](https://github.com/traefik/yaegi)             |            1000000            |              1184 ns/op | interpreter |
| [pkujhd/goloader](https://github.com/pkujhd/goloader)         |           68201743            |             19.11 ns/op |   native    |
| [tetratelabs/wazero](https://github.com/tetratelabs/wazero)   |           11401358            |             105.0 ns/op |   native    |
 
Several of the other packages use RPC or similar methods instead of the go plugin package which gets around issues such as, but not limited to, [not being compatible with Windows](https://github.com/golang/go/issues/19282) and [package paths and GOPATH needing to be the same between apps and plugins](https://github.com/golang/go/issues/20481).

With the addition of [Yaegi](https://github.com/traefik/yaegi), I am also benchmarking interpreters. 

**Do you know any other plugin packages?** Please open an [issue](https://github.com/uberswe/go-plugin-benchmark/issues/new) or pull request.

**Found an issue with a benchmark?** Please open an [issue](https://github.com/uberswe/go-plugin-benchmark/issues/new) or pull request.

## Setup

I have moved to using [Docker](https://www.docker.com/) to run these benchmarks to make them easier to reproduce.

Build the docker image:

```
docker build --pull --no-cache --tag go-plugin-benchmark:local .
```

To run the benchmark run

```
docker run go-plugin-benchmark:local
```

## Results

Most plugins tested are using RPC which adds about 30 - 50 microseconds to plugin calls (or 0.03 - 0.05 milliseconds) over the golang plugin package.

The [goloader](https://github.com/pkujhd/goloader) package is interesting and may provide a good alternative to the go plugin package. One drawback is that it uses internal packages which requires renaming the internal folder locally and I have not tested compatibility to see if it solves the problems with the go plugin package.

## Contributing

If you have a plugin or interpreter written in go which you would like to benchmark feel free to [open an issue](https://github.com/uberswe/go-plugin-benchmark/issues/new).

Would you like to add more benchmarks? Please feel free to fork this repository and open a pull request with the updated changes. Please make sure to add any needed code and also update the Readme. Your code belongs to you but it must fall under the same LICENSE as this repository to be included.