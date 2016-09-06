package main

import (
	"fmt"
	_ "math"
	"mymodule"
)

func main() {
	foo()
}

func foo() {
	fmt.Println("foo() was called!")
	calculate(1.0, 2.0, 3.0)
	res1, res2, resName := compute(2.0, 3.0, 4.0, "compute")
	fmt.Println(res1, res2, resName)

	fmt.Println(add(1, 2, 3))
	fmt.Println(add(2, 3, 4, 5, 6))
	closure_func("closure function")
	fmt.Println(fibonacci(10))

	fmt.Println("=================================================")

	fmt.Println("access	mymodule…")
	var	c	int
	c	=	mymodule.Add(10,6)
	fmt.Printf("add()=%d\n",c)
	c	=	mymodule.Subtract(5,8)
	fmt.Printf("subtract()=%d\n",c)
	c	=	mymodule.Multiply(5,3)
	fmt.Printf("multiply()=%d\n",c)

}

func calculate(a, b, c float64) float64 {
	result := a * b * c
	fmt.Printf("a=%.2f,	b=%.2f,	c=%.2f	=	%.2f	\n", a, b, c, result)
	return result
}

//	a	function	with	parameters	and	multiple	return	values
func compute(a, b, c    float64, name    string) (float64, float64, string) {
	result1 := a * b * c
	result2 := a + b + c
	result3 := result1 + result2
	newName := fmt.Sprintf("%s	value	=	%.2f", name, result3)
	return result1, result2, newName
}


//	a	function	with	zero	or	more	parameters	and	a	return	value
func add(numbers    ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}



//	a	closure	function
func closure_func(name    string) {
	hoo := func(a, b    int) {
		result := a * b
		fmt.Printf("hoo()	=	%d	\n", result)
	}
	joo := func(a, b    int) int {
		return a * b + a
	}
	fmt.Printf("closure_func(%s)	was	called\n", name)
	hoo(2, 3)
	val := joo(5, 8)
	fmt.Printf("val	from	joo()	=	%d	\n", val)
}


//	a	recursion	function
func fibonacci(n    int) int {
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}
	return (fibonacci(n - 1) + fibonacci(n - 2))
}