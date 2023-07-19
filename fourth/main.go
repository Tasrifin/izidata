package main

import "fmt"

func main() {
	concatString("JOHN", "SMITH")
}

func concatString(first, second string) {
	var result string
	maxLen := len(first)
	if len(second) > maxLen {
		maxLen = len(second)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(first) {
			result = result + string(first[i])
		}
		if i < len(second) {
			result = result + string(second[i])
		}
	}

	fmt.Printf("%s + %s = %s", first, second, result)
}
