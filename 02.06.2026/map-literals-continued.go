package main

import "fmt"

type Rertex struct {
	Lat, Long float64
}

var b = map[string]Rertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main4() {
	fmt.Println(b)
}
