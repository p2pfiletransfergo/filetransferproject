package main

import (
	"encoding/binary"
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

	
	binary.Write(conn, binary.BigEndian, uint32(len(p)))
	conn.Write([]byte(p)) 
	fileinfo, _ := file.Stat()
	filesize := fileinfo.Size()
	binary.Write(conn, binary.BigEndian, filesize)
	io.Copy(conn, file)

	// conn.Write([]byte(p))
	// io.Copy(conn, file)
}
