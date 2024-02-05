package main

import (
	"assign2prob1/palindrome"
	"fmt"
)

func main() {
	var name string = "malayalam"
	res := palindrome.Test(name)
	fmt.Println(res)
}
