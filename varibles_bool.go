package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println("Hello Naveen")
	fmt.Println(x, y, z)
	z = true
	if z {
		fmt.Println("true")
	}
	func1()
}

func func1() {
	a := 10
	b := 20
	fmt.Println(a == b)
	fmt.Println(a < b)
	fmt.Println(a > b)

}
