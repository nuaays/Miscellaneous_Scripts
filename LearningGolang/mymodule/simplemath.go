package mymodule

import "fmt"

func init(){
	fmt.Println("simplemath ... ...")
}

func Add(a, b int) int {
	return a + b
}
func Subtract(a, b int) int {
	return a - b
}
func Multiply(a, b int) int {
	return a * b
}