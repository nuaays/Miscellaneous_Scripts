package main

import (
	"fmt"
	"encoding/json"

)

func main() {
	fmt.Println("Demo encoding JSON")
	type Employee struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
	}


	//struct to json
	fmt.Println(">>>>>struct to json ...")
	emp := &Employee{Id:"12345", Name:"test", Email:"test@gmail.com"}
	b, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))


	//json to struct
	fmt.Println(">>>>>json to struct")
	var newEmp Employee
	str := `{"Id":"4566","Name":"Brown","Email":"brown@email.com"}`
	json.Unmarshal([]byte(str), &newEmp)
	fmt.Printf("Id   : %s\n", newEmp.Id)
	fmt.Printf("Name : %s\n", newEmp.Name)
	fmt.Printf("Email: %s\n", newEmp.Email)

}
