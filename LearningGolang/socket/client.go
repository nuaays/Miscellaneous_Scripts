package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SVR_HOST = "localhost"
	SVR_PORT = "9982"
	SVR_TYPE = "tcp"
)

func main() {
	fmt.Println(os.Args[1])
	fmt.Println("client	is	running")
	conn, err := net.Dial(SVR_TYPE, SVR_HOST + ":" + SVR_PORT)
	if err != nil {
		panic(err)
	}
	fmt.Println("send data")
	_, err = conn.Write([]byte("message	from client:" + os.Args[1]))
	defer conn.Close()
}