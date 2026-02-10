package main

import (
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

func sendtcp(addr string, path string) {
	//addr := <-ip
	time.Sleep(2 * time.Second)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// fmt.Print("Enter file path: ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// path := scanner.Text()

	//fmt.Println("You chose:", path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := file.Name()
	p = filepath.Base(p)
	conn.Write([]byte(p))
	io.Copy(conn, file)
}
