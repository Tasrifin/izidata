package main

import "fmt"

func main() {
	checkYear(2016)
}

func checkYear(year int) {
	result := false

	if year%4 == 0 {
		result = true
	}
	fmt.Printf("%v is LEAP? %v", year, result)
}
