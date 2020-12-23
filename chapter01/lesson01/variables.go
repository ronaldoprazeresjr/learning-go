package main

import "fmt"

var z = 50 // package scope

func main() {
	x := 10 // scope block
	fmt.Println(x)
	x = 20
	fmt.Println(x)
	x, y := 30, 40
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
