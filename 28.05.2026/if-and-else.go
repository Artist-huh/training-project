package main

import (
	"fmt"
	"math"
)

func low(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func el() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
