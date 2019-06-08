package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(a int, b int) int {
	return a + b
}

func fibonacci(n int) int {
	//0 1 1 2 3 5 8 13 21 etc
	if n == 0 {
		return 0
	}

	var result int = 1
	var previous int = 0
	var temp int = 0
	for i := 1; i < n; i++ {
		temp = result
		result = sum(result, previous)
		previous = temp
	}
	return result
}

func fibonacci_recursive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci_recursive(n-1) + fibonacci_recursive(n-2)
	}
}

func is_prime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//multiple return example
func sum_and_diff(a int, b int) (int, int) {
	return sum(a, b), a - b
}

func prime(n int) int {
	count := 0
	prime := 0
	for i := 2; count != n; i++ {
		if is_prime(i) {
			count++
			prime = i
		}
	}
	return prime
}

func main() {
	fmt.Printf("3 + 5 = %d\n", sum(3, 5))
	fib := 8
	// fmt.Printf("Fibonacci_recursive(%d) = %d\n", fib, fibonacci_recursive(fib))
	if len(os.Args) > 1 {
		fib, _ = strconv.Atoi(os.Args[1])
		fmt.Printf("Fibonacci(%d) = %d\n", fib, fibonacci(fib))
	}

	fmt.Printf("The %d prime number is %d\n", 100, prime(100))
}
