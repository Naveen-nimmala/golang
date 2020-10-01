package main

import (
	"fmt"
)

func main() {
	if i := 10; i == 1 {
		fmt.Println("True 01")
	} else if i == 5 {
		fmt.Println("5")
	} else {
		fmt.Println("this is not true")
	}
}
