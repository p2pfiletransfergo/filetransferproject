package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func handleconn(conn net.Conn) {
	defer conn.Close()

	var nameLen uint32
	binary.Read(conn, binary.BigEndian, &nameLen)

	nameBuf := make([]byte, nameLen)
	io.ReadFull(conn, nameBuf)
	filename := string(nameBuf)

	var fileSize int64
	binary.Read(conn, binary.BigEndian, &fileSize)
	original := filename

	for i := 0; ; i++ {

		_, err := os.Stat("rec_" + filename)

		if os.IsNotExist(err) {
			break
		}

		filename = strconv.Itoa(i) + "_" + original
	}
	file, _ := os.Create("rec_" + filename)
	defer file.Close()

	io.CopyN(file, conn, fileSize)

	fmt.Println("Received:", filename)
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
			fmt.Println("connection closed")
			return
		}
		go handleconn(conn)
	}
}
