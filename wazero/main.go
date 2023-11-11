package main

import (
	"math/rand"
)

func main() {}

//export RandInt
func RandInt() int {
	return rand.Int()
}
