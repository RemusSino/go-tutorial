package main

import "fmt"

func arrays() {
	//arrays
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
}

func slices() {
	shapeSlice := []string{"Circle", "Triangle"}
	//add an element
	shapeSlice = append(shapeSlice, "Square")
	//length of slice
	fmt.Printf("I have %d shapes", len(shapeSlice))
	fmt.Println("I draw this shapes ", shapeSlice)

	friendsSlice := make([]string, 0)
	friendsSlice = append(friendsSlice, "Bob", "John", "Marcus", "Helen")
	fmt.Println("My friends are ", friendsSlice)

	maleFriendsSlice := friendsSlice[0:3]
	fmt.Println("I'm playing football with ", maleFriendsSlice)
}

func maps() {

}

func main() {
	arrays()
	slices()
}
