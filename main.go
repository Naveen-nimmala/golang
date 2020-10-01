package main

import (
	"fmt"
)

func main() {
	x := 10
	var y int
	var z int = 20
	fmt.Printf("%T\n%T\n%T\n", x,y,z)
	fmt.Println(x,y,z)
	hello()
}

func hello(){
	fmt.Println("this is from Hello")
}
