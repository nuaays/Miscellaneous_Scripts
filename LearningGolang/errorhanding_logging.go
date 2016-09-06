package main

import (
	"fmt"
	"log"
	_ "os"
	"os"
)

func simpleLogging() {
	fmt.Println("===== simple logging =====")
	log.Println("Hello World")
	log.Println("This is a simple error")

}

func formatLogging() {
	fmt.Println("===== formatting Logging =====")
	var warning *log.Logger

	warning = log.New(
		os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning.Println("This is warning message 1")
	warning.Println("This is warning message 2")


}

func fileLogging() {
	fmt.Println("===== file logging =====")
	file, err := os.OpenFile("d:/myapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file")
		return
	}

	var logFile *log.Logger
	logFile = log.New(
		file,
		"APP: ",
		log.Ldate|log.Ltime|log.Lshortfile	)

	logFile.Println("This is error message 1")
	logFile.Println("This is error message 2")
	logFile.Println("This is error message 3")
	fmt.Println("Done")
}

func main(){
	//calculate()
	simpleLogging()
	formatLogging()
	fileLogging()




}





func calculate()	{
	fmt.Println("----demo	error	handling---")
	//a	:=	10
	//b	:=	2
	//c	:=	1.0
	//c	=	float64(a/b)
	c := 0

	defer func() {
		a := 10
		b := 0
		c =a/b
		if error := recover(); error != nil {
			fmt.Println("Recovering .. ", error)
			fmt.Println(error)
		}
	}()

	fmt.Printf("result	=	%.2f	\n",c)
	fmt.Println("Done")
}