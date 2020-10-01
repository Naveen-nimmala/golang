package main

import (
	"fmt"
)

const (
	x = 10
	y = 20.123
	c = "hello"
)

func main() {
	fmt.Println(x, y, c)
	fmt.Printf("%d%T\n", x, x)
	fmt.Printf("%d%T\n", y, y)
	fmt.Printf("%d%T\n", c, c)

}
