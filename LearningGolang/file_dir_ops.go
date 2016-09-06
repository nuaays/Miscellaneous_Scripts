package main

import (
	"fmt"
	"io/ioutil"
)


func writeFile(message	string)	{
	bytes	:=	[]byte(message)
	ioutil.WriteFile("d:/testgo.txt",bytes,0644)
	fmt.Println("created	a	file")
}

func readFile() {
	data, _ := ioutil.ReadFile("d:/testgo.txt")
	fmt.Println("File Content:")
	fmt.Println(string(data))
}

func main() {
	writeFile("hello world")
	readFile()


}
