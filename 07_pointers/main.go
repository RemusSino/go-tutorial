package main

import "fmt"

func change(x *int) {
	*x = 100
}

func notChange(x int) {
	x = 101
}

//go is pass by value
func main() {
	x := 2
	change(&x)
	fmt.Println(x)
	notChange(x)
	fmt.Println(x)
}
