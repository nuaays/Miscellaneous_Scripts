package main

import (
	"fmt"
	"runtime"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("hello!")
	tet := fmt.Sprintf("OS:%s, Arch:%s", runtime.GOOS, runtime.GOARCH);
	fmt.Println(tet)
	fmt.Println()


	a	:=	2.4
	b	:=	1.6
	c	:=	math.Pow(a,2)
	fmt.Printf("%.2f^%d	=	%.2f	\n",a,2,c)
	c	=	math.Sin(a)
	fmt.Printf("Sin(%.2f)	=	%.2f	\n",a,c)
	c	=	math.Cos(b)
	fmt.Printf("Cos(%.2f)	=	%.2f	\n",b,c)
	c	=	math.Sqrt(a*b)
	fmt.Printf("Sqrt(%.2f*%.2f)	=	%.2f	\n",a,b,c)

	d	:=	4
	//	increment
	fmt.Printf("a	=	%d	\n",d)
	d	=	d	+	1
	fmt.Printf("a	+	1	=	%d	\n",d)
	d++
	fmt.Printf("a++	=	%d	\n",d)
	//	decrement
	d	=	d	-	1
	fmt.Printf("a	-	1	=	%d	\n",d)
	d--
	fmt.Printf("d--	=	%d	\n",d)


	var	radius	float64 = 20.0
	//fmt.Print("Enter	a	radius	value:	")
	//fmt.Scanf("%f",	&radius)
	area	:=	math.Pi	*	math.Pow(radius,2)
	fmt.Printf("Circle	area	with	radius	%.2f	=	%.2f	\n",radius,	area)

	e := 5
	f := 8
	if e > f || e -f < e {
		fmt.Println("Conditional -->a>b || a-b<a")
	}else{
		fmt.Println("...another")
	}

	selected	:=	2
	switch	selected	{
	case	0:
		fmt.Println("selected	=	0")
	case	1:
		fmt.Println("selected	=	1")
	case	2:
		fmt.Println("selected	=	2")
	case	3:
		fmt.Println("selected	=	3")
	default:
		fmt.Println("other..")
	}


	//
	var	i	int
	for	i=0;i<5;i++	{
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	for	j:=5;j<11;j++	{
		if j==7 {
			continue
		}
		fmt.Println(j)
	}

	//
	var	numbers[5]	int
	var	cities[5]	string
	var	matrix[3][3]	int	//	array	2D

	fmt.Println(">>>>>insert	array	data")
	for	i:=0;i<5;i++	{
		numbers[i]	=	i
		cities[i]	=	fmt.Sprintf("City	%d",i)
	}
	fmt.Println(">>>>>display	array	data")
	for	j:=0;j<5;j++	{
		//fmt.Println(numbers[j])
		fmt.Println(cities[j])
	}
	//
	fmt.Println(">>>>>insert	matrix	data")
	for	i:=0;i<3;i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	fmt.Println(">>>>>display	matrix	data")
	for	i:=0;i<3;i++	{
		for	j:=0;j<3;j++	{
			fmt.Println(matrix[i][j])
		}
	}


	//slice
	//	define	slice
	fmt.Println("define	slices")
	var	numbers1[]	int
	numbers1	=	make([]int,5)
	matrix1	:=	make([][]int,3*3)
	//	insert	data
	fmt.Println(">>>>>insert	slice	data")
	for	i:=0;i<5;i++	{
		numbers1[i]	=	rand.Intn(100)	//	random	number
	}
	//	insert	matrix	data
	fmt.Println(">>>>>insert	slice	matrix	data")
	for	i:=0;i<3;i++	{
		matrix1[i]	=	make([]int,3)
		for	j:=0;j<3;j++	{
			matrix1[i][j]	=	rand.Intn(100)
		}
	}
	//	display	data
	fmt.Println(">>>>>display	sclice	data")
	for	j:=0;j<len(numbers1);j++	{
		fmt.Println(numbers1[j])
	}
	//	display	matrix	data
	fmt.Println(">>>>>display	sclice	2D	data")
	for	i:=0;i<3;i++	{
		for	j:=0;j<3;j++	{
			fmt.Println(matrix1[i][j])
		}
	}



	//	define	array
	fmt.Println("define	map")
	products	:=	make(map[string]int)
	customers	:=	make(map[string]int)
	//	insert	data
	fmt.Println(">>>>>insert	map	data")
	products["product1"]	=	rand.Intn(100)
	products["product2"]	=	rand.Intn(100)
	customers["cust1"]	=	rand.Intn(100)
	customers["cust2"]	=	rand.Intn(100)
	//	display	data
	fmt.Println(">>>>>display	map	data")
	fmt.Println(products["product1"])
	fmt.Println(products["product2"])
	fmt.Println(customers["cust1"])
	fmt.Println(customers["cust2"])


}