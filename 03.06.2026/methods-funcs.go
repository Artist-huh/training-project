package main

import (
	"fmt"
	"math"
)

type Lertex struct {
	X, Y float64
}

func Abs(v Lertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main52() {
	v := Lertex{3, 4}
	fmt.Println(Abs(v))
}
