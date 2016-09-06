package main

import "fmt"

func main() {
	fmt.Println("=== simple channels===")

	//define a channel
	c := make(chan int)

	//run a function in backgroud
	go func() {
		fmt.Println("goroutine process")
		c <- 10 //write data to a channel
	}()

	val := <- c //read data from a channel
	fmt.Printf("value: %d\n", val)


}
