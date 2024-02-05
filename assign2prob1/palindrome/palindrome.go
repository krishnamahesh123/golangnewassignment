package palindrome

import (
	"strings"
)

func Test(st string) string {
	st_convert := strings.ToLower(st)
	var reverse = ""
	for i := len(st_convert) - 1; i >= 0; i-- {
		reverse = reverse + string(st_convert[i])
	}
	if st_convert == reverse {
		return "given string is palindrome"
	}
	return "given string is not a palindrome"
}
