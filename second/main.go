package main

import "fmt"

func main() {
	printFibonacci(8)
}

func printFibonacci(n int) {
	a := 0
	b := 1
	c := 0

	for i := 1; i <= n; i++ {
		fmt.Printf("%v ", a)
		c = a + b
		a = b
		b = c
	}
}
