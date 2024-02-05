package main

import (
	"assign2prob3/pack6"
	"fmt"
)

func main() {
	var sl = []int{10, 20, 30, 40, 50}
	key := 40
	res := pack6.BinarySearch(sl, key)
	if res == -1 {
		fmt.Println("Key is not found")
	} else {
		fmt.Println("Key is found")
	}
}
