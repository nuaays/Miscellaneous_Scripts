package main


import (
	"fmt"
	"time"
)


func calculate(){
	i := 12
	for i < 15 {
		fmt.Printf("calculate() = %d \n", i)
		var result float64 = 8.2 * float64(i)
		fmt.Printf("calculate() result=%2.f\n", result)
		time.Sleep(500 * time.Millisecond)
		i++
	}
}

func main() {
	fmt.Println("goroutines demo\n")

	//run func in backgroud
	go calculate()

	index := 0
	//run in backgroud
	go func() {
		for index < 6 {
			fmt.Printf("go func() index=%d \n", index)
			var result float64 = 2.5 * float64(index)
			fmt.Printf("go func() result=%2.f\n", result)

			time.Sleep(500 * time.Millisecond)
			index++
		}
	}()

	//run in backgroud
	fmt.Println("======================")
	var input string
	fmt.Scan(&input)
	fmt.Println("Done")

}
