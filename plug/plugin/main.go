package main

import (
	"fmt"
	"github.com/uberswe/go-plugin-benchmark/plug"
	"math/rand"
)

type RandomIntService struct{}

func (ris *RandomIntService) Get(in int64) (int64, error) {
	return int64(rand.Int()), nil
}

func main() {
	err := plug.RunRandomIntService(&RandomIntService{})
	if err != nil {
		fmt.Println(err.Error(), ", exiting...")
	}
}
