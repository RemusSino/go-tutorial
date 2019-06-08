package main

import (
	"fmt"
	"strconv"
	"strings"
)

var stringBuilder strings.Builder

type Car struct {
	carType string
	color   string
	seats   int
}

type Person struct {
	Car
	name    string
	age     int
	country string
}

func (p Person) myCar() string {
	stringBuilder.Reset()
	stringBuilder.WriteString("My car is ")
	stringBuilder.WriteString(p.color)
	return stringBuilder.String()
}

func (p Person) toString() string {
	stringBuilder.Reset()
	stringBuilder.WriteString("My name is ")
	stringBuilder.WriteString(p.name)
	stringBuilder.WriteString(", I am ")
	stringBuilder.WriteString(strconv.Itoa(p.age))

	return stringBuilder.String()
}

func main() {
	famillyCar := Car{"Van", "Blue", 7}
	me := Person{famillyCar, "Remus", 30, "RO"}

	fmt.Println(me.toString())
	fmt.Println(me.myCar())
}
