package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"time"
	"fmt"
)

func main() {
	fh, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		os.Exit(1)
	}
	defer fh.Close()

	//logging
	l := log.New(fh, "", os.O_APPEND)

	cmd := exec.Command("ls", "-l")
	fmt.Println(cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Start()
	r := bufio.NewReader(stdout)
	line, _, err := r.ReadLine()
        fmt.Println(string(line))
	l.Printf("[%s] %s\n", time.Now().Format("2006-01-02 12:12:12"), string(line))
	time.Sleep(time.Second * 1)

}
