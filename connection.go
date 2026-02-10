package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleconn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("recieved file name :", n, string(buf[:n]))
	file, err := os.Create("rec_" + string(buf[:n]))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	n1, err := io.Copy(file, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Recieved bytes :", n1)
}
func listentcp() {
	ln, err := net.Listen("tcp", ":5050")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			//panic(err)
			fmt.Println("connection closed")
			return
		}
		go handleconn(conn)
	}
}
