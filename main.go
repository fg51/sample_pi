package main

import (
	"fmt"
	"math/rand"
	"time"
	"./gencmplx"
)


func main() {
	c := gencmplx.New(rand.NewSource(time.Now().UnixNano()), int64(10000))
	for i := range c {
		fmt.Printf("%v\t%v\n", real(i), imag(i))
	}
}
