package main

import (
	"fmt"
	"math"
)

type Rertex struct {
	X, Y float64
}

func (v Rertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Rertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Rertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
