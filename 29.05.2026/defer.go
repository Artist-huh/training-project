package main

import "fmt"

func m() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
