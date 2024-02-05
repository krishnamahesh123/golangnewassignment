package main

import (
	"assign2prob2/pack5"
	"fmt"
)

func main() {
	str1 := "Listen"
	str2 := "siLent"
	res := pack5.Check(str1, str2)
	if res {
		fmt.Println("The given strings are Anagrams")
	} else {
		fmt.Println("The given strings are not Anagrams")
	}
}
