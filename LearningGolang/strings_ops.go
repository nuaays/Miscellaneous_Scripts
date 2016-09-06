package main

import (
	"fmt"
	_ "strconv"
	_ "strings"
	"strings"
)

func main() {
	str1 := "hello"
	str2 := "world"
	str3 := str1 + str2
	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str3)
	fmt.Println(&str1, &str2, &str3, &str4)


	//split
	data := "Berlin;Amsterdam;London;Tokyo"
	cities := strings.Split(data, ";")
	for _, city := range cities {
		fmt.Println(city)
	}
	//slice
	mystr := "AbcDefg"
	fmt.Println(mystr[2:5])
	fmt.Println(mystr[len(mystr)-3:len(mystr)])
	//Upper, Lower
	fmt.Println(strings.ToUpper(mystr))
	fmt.Println(strings.ToLower(mystr))
	fmt.Println(mystr)


}
