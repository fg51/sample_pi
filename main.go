package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"time"

	"./gencmplx"
)


func main() {
	c := gencmplx.New(rand.NewSource(time.Now().UnixNano()), int64(10000))
	total := int64(0)
	in_circle := int64(0)
	for i := range c {
		total++
		if cmplx.Abs(i) <= float64(1) {
			in_circle++
		}
	}
	fmt.Printf("total = %v, hit = %v, pi = %v \n", total, in_circle, float64(4 * in_circle) / float64(total))
}
