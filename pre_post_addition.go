package main

import "fmt"

func main() {
	x := 7
	y := 3

	x = y
	y++
	fmt.Printf("The value of x = y++ is %d\n", x)

	y++
	x = y
	fmt.Printf("The value of x = ++y is %d\n", x)
}