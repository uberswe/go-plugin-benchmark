#!/usr/bin/env sh

protoc -I=. -I=$(go list -f '{{.Dir}}' github.com/elliotmr/plug/pkg/plugpb) --go_out=. --go_opt=module=uberswe/go-plugin-benchmark/plug --plug_out=. --plug_opt=module=uberswe/go-plugin-benchmark/plug plugin.proto
