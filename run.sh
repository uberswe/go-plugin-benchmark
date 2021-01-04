#!/bin/bash
go build -buildmode=plugin -o plugin.so golangplugin/main.go
go build -o ./hashicorpgoplugin ./hashicorp-go-plugin/main.go
go build -o ./pieplugin ./pie/main.go
go build -o ./pingoplugin ./pingo/main.go
go test -bench=.