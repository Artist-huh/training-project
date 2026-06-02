package main

import "fmt"

type Lertex struct {
	Lat, Long float64
}

var a = map[string]Lertex{
	"Bell Labs": Lertex{
		40.68433, -74.39967,
	},
	"Google": Lertex{
		37.42202, -122.08408,
	},
}

func main3() {
	fmt.Println(a)
}
