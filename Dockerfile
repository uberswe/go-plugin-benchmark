FROM golang:1.21

WORKDIR /app

ENV GO111MODULE=on
ENV GOARCH=amd64

RUN apt-get update
RUN apt-get install gcc
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.30.0/tinygo_0.30.0_amd64.deb && dpkg -i tinygo_0.30.0_amd64.deb
RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.13.4

COPY . ./

RUN go get
RUN cp -r /usr/local/go/src/cmd/internal /usr/local/go/src/cmd/objfile
RUN go build -buildmode=plugin -o plugin.so golangplugin/main.go
RUN go build -o ./hashicorpgoplugin ./hashicorp-go-plugin/main.go
RUN go build -o ./pieplugin ./pie/main.go
RUN go build -o ./pingoplugin ./pingo/main.go
RUN go build -o ./plugplugin ./plug/plugin/main.go
RUN tinygo build -o ./wasm-edge.wasm -target wasi ./wasm-edge/main.go
RUN go list -export -f '{{if .Export}}packagefile {{.ImportPath}}={{.Export}}{{end}}' std `go list -f {{.Imports}} ./goloader/main.go | awk '{sub(/^\[/, ""); print }' | awk '{sub(/\]$/, ""); print }'` > importcfg
RUN CGO_ENABLED=0 go tool compile -importcfg importcfg -o ./goloader.o ./goloader/main.go

CMD . $HOME/.wasmedge/env && go test -bench=.
