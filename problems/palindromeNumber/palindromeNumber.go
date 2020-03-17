package palindromeNumber

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	xstr:=strings.Split(strconv.Itoa(x), "")
	for i,j:=0,len(xstr)-1;i<j; {
		if xstr[i] !=xstr[j]{
			return false
		}
		i++
		j--
	}
	return true
}