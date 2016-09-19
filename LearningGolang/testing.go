package main

import "fmt"

func main() {
	array := [...]string{"Go", "Python", "Java", "C", "C++", "PHP", "Ruby"}
	fmt.Println(len(array))
	slice := array[3:]
	fmt.Println(cap(slice))
	slice1 := append(slice, "Perl", "Erlang")
	fmt.Println(slice1)

	sliceA := []string{"Notepad", "Ultraedit", "Eclipse"}
	sliceB := []string{"Vim", "Emacs", "LiteIDE", "IDEA"}

	n1 := copy(sliceA, sliceB)
	fmt.Println(n1)

}
