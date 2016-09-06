package main

import (
	"fmt"
	"time"
	"math"
)

//define a struct
type Employee struct {
	id int
	name string
	country string
	//created time.Time
	created string
}

//define a circle
type Circle struct {
	x, y int
	r float64
}

func (c Circle) display() {
	fmt.Printf("x=%d, y=%d, r=%f\n", c.x, c.y, c.r)
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c Circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

func (c Circle) moveTo(newX, newY int) {
	c.x, c.y = newX, newY
}


func main() {
	//
	shape := Circle{10, 5, 2.8}
	shape.display()
	fmt.Printf("area=%2.f\n", shape.area())
	fmt.Printf("circumference=%2.f\n", shape.circumference())

	shape.moveTo(5,5)
	shape.display()

	//
	var emp Employee
	newEmp := new(Employee)

	//
	emp.id = 2
	emp.name = "Employee 2"
	emp.country = "DE"
	emp.created = time.Now().String()

	newEmp.id = 5
	newEmp.name = "Employee 5"
	newEmp.country = "UK"
	newEmp.created  = time.Now().String()


	fmt.Println("=======================")
	fmt.Println(emp)
	fmt.Println(*newEmp)



}
