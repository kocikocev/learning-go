package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary int
}

type Point struct {
	x int
	y int
}

var me Employee

func main() {
	var employeeOfTheMonth *Employee = &me

	employeeOfTheMonth.Name += "Best"
	fmt.Println(employeeOfTheMonth.Name)

	(*employeeOfTheMonth).Name += "best"
	fmt.Println(me.Name)

	// Structs are commonly dealt through pointers, the following are equivalent:
	pp := &Point{1, 2}
	pp = new(Point)
	*pp = Point{1, 2}
}
