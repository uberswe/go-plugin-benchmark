#!/usr/bin/env sh

protoc -I=. --go_out=. --go_opt=module=uberswe/go-plugin-benchmark/plug --plug_out=. --plug_opt=module=uberswe/go-plugin-benchmark/plug plugin.proto
