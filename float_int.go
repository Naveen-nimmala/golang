package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

// byte is eques to uint8
// rune is eqls to int32
func main() {
	x = 10
	y = 10.123
	fmt.Printf("%T\n%T\n", x, y)
	fmt.Println(runtime.GOOS)
	//fmt.Println(runtime.MemStats)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Version)
	fmt.Println(runtime.NumCPU)
}
