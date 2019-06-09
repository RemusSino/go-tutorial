package main

import (
	"fmt"
	"sync"
)

// WaitGroup
// Mutex
// runtime.Gosched()
// package atomic
var wg sync.WaitGroup

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
		result = result + previous
		previous = temp
	}

	return result
}

func fibonacciWg(n int) {
	fmt.Printf("Fibonacci(%d)= %d\n", n, fibonacci(n))
	wg.Done()
}

func fibonacciChan(n int, channel chan<- (int)) {
	channel <- fibonacci(n)
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}

func factorialChan(n int, channel chan<- (int)) {
	channel <- factorial(n)
}

func main() {
	// fmt.Print(factorial(3))
	wg.Add(1)
	fmt.Println("Calculating Fibonacci(100) in goroutine with WaitGroup")
	go fibonacciWg(100)
	fmt.Println("Waiting for goroutine to finish...")
	wg.Wait()

	fibChannel := make(chan int)
	fmt.Println("Calculating Fibonacci(100) in goroutine with Channel")
	go fibonacciChan(100, fibChannel)
	factChannel := make(chan int)
	fmt.Println("Calculating Factorial(10) in goroutine with Channel")
	go factorialChan(10, factChannel)

	for i := 0; i < 2; i++ {
		select {
		case fib := <-fibChannel:
			fmt.Println("Fibonnaci(100)=", fib)
		case fact := <-factChannel:
			fmt.Println("Factorial(10)=", fact)
		}
	}
}
