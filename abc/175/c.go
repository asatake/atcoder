package main

import (
	"fmt"
	"math"
)

func main() {
	var x, k, d float64
	fmt.Scan(&x, &k, &d)

	x = math.Abs(x)

	st := math.Min(k, float64(int64(x) / int64(d)))
	k -= st
	x -= st * d

	if int64(k) % 2 == 0 {
		fmt.Printf("%.0f\n", x)
	} else {
		fmt.Printf("%.0f\n", d - x)
	}
}
