package main

import (
	"fmt"
	"strings"
)

func main() {
	checkPalindrome("rttaroo")
}

func checkPalindrome(char string) {
	checkData := make(map[string]int)
	countCheck := 0
	result := true

	for _, v := range strings.ToLower(char) {
		checkData[string(v)]++
	}

	for _, total := range checkData {
		if total%2 != 0 {
			countCheck++
		}
	}

	if countCheck > 1 {
		result = false
	}

	fmt.Printf("Check is %v palindrome : %v", char, result)
}
