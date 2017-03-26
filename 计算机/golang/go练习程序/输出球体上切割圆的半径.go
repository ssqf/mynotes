package main

import (
	"fmt"
	"math"
)

func main() {
	rn := 46.0
	detal := 0.0
	preR := rn
	curR := rn
	alpha := 90.0 / 32.0 * math.Pi / 180
	fmt.Printf("alpha=%f\n", alpha)
	for i := 0; i <= 32; i++ {
		curR = rn * math.Cos(alpha*float64(i))
		detal = preR - curR
		preR = curR
		fmt.Printf("r(%d) = %f %f\n", i, curR, detal)
	}
}
