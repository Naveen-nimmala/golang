package main

import (
	"fmt"
)

func main() {
	str := "Hello work"
	srt1 := `"hello Naveen
How are you
bye"`
	fmt.Println(str)
	fmt.Println(srt1)
	bytes()
}

func bytes() {
	s := "Hello Naveen"
	convertByte := []byte(s)
	fmt.Println(convertByte)
	fmt.Printf("%T", convertByte)
}
