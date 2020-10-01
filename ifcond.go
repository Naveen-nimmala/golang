package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("True 01")
	}
	if false {
		fmt.Println("False 02")
	}
	if !(true) {
		fmt.Println("03")
	}
	if !(false) {
		fmt.Println("04")
	}
	cond()
}

func cond() {
	if 2 == 2 {
		fmt.Println("01")
	}
	if 2 != 2 {
		fmt.Println("02")
	}
	if !(2 == 2) {
		fmt.Println("03")
	}
	if !(2 != 2) {
		fmt.Println("04")
	}
	if x := 10; x == 10 {
		fmt.Println(x)
	}
}
