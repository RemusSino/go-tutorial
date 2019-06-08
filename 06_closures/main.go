package main

import "fmt"

type fn func(int, int) int

func printFuncResult(f fn, a int, b int) {
	fmt.Printf("Result is %d \n", f(a, b))
}

func main() {
	printFuncResult(func(a int, b int) int {
		return a + b
	}, 3, 5)

	printFuncResult(func(a int, b int) int {
		return a * b
	}, 3, 5)
}
