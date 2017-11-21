package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		c := complex(
			float64(rand.Int63n(10000001)) / float64(10000000),
			float64(rand.Int63n(10000001)) / float64(10000000))
		fmt.Printf("%v (%v)\n", c, cmplx.Abs(c))
	}
}
