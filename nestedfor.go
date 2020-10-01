package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("This is from first", i)
		for j := 0; j < 3; j++ {
			fmt.Println("\tTHis is second loop", j)
		}
	}
}
