package pack5

import "fmt"

func Check(str1 string, str2 string) bool {
	map1 := make(map[string]int)
	if len(str1) != len(str2) {
		return false
	}
	for _, ch := range str1 {
		map1[string(ch)]++
	}
	fmt.Println(map1)
	for _, ch := range str2 {
		map1[string(ch)] = map1[string(ch)] - 1
		if map1[string(ch)] < 0 {
			fmt.Println("check")
			return false

		}

	}
	fmt.Println(map1)
	return true

}
