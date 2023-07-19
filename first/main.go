package main

import "fmt"

func main() {
	printPyramid(8)
}

func printPyramid(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}
}
