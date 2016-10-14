package main

import (
	"fmt"
	"github.com/xeonx/timeago"
	"time"
)

func main() {
	t := time.Now().Add(42 * time.Second)

	s := timeago.English.Format(t)
	//s will contains "less than a minute ago"
	fmt.Println(s)

	//...
}
