package main

import "fmt"

func conditional() {
	var a int
	a = 3
	b := 22

	if a < b {
		fmt.Printf("%d < %d\n", a, b)
	}
}

func forLoop() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Int values are")
	for i := 0; i < 10; i++ {
		fmt.Println(intSlice[i])
	}

	fmt.Println("Printing indexes and values")
	for index, value := range intSlice {
		fmt.Printf("%d : %d\n", index, value)
	}

	fmt.Println("Printing indexes and values, for odd values")
	for index, value := range intSlice {
		if value%2 != 0 {
			fmt.Printf("%d : %d\n", index, value)
		}
	}
}

func main() {
	conditional()
	forLoop()

}
