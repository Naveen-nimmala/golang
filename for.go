package main

import (
	"fmt"
)

func main() {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("Hello Naveen")
		i++
	}

	for i < 20 {
		fmt.Println("Hello Naveen", i)
		i++
	}
	for j := 0; j < 10; j++ {
		fmt.Println("Hello Naveen")
	}
}
