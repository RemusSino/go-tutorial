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

func fibonacciChan(n int, channel chan (int)) {
	channel <- fibonacci(n)
}

func main() {
	wg.Add(1)
	fmt.Println("Calculating Fibonacci(100) in goroutine with WaitGroup")
	go fibonacciWg(100)
	fmt.Println("Waiting for goroutine to finish...")
	wg.Wait()

	channel := make(chan int)
	fmt.Println("Calculating Fibonacci(100) in goroutine with Channel")
	go fibonacciChan(100, channel)
	fmt.Println("Fibonnaci(100)=", <-channel)

}
