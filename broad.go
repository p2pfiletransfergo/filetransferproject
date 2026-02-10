package main

import (
	"net"
	"time"
)

func getIP() *net.UDPAddr {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr
}

//	func brodudp() {
//		conn, err := net.Dial("udp4", "10.126.7.255:8080")
//		if err != nil {
//			panic(err)
//		}
//		defer conn.Close()
//		for {
//			p := getIP()
//			_, err1 := conn.Write([]byte(p.IP.String() + ":5050"))
//			if err1 != nil {
//				panic(err1)
//			}
//			time.Sleep(1 * time.Second)
//		}
//	}
func brodudp() {
	//	raddr, _ := net.ResolveUDPAddr("udp4", "10.132.122.255:8080")
	raddr, _ := net.ResolveUDPAddr("udp4", "255.255.255.255:8080")

	conn, _ := net.DialUDP("udp4", nil, raddr)
	conn.SetWriteBuffer(1024)

	for {
		p := getIP()
		conn.Write([]byte(p.IP.String() + ":5050"))
		time.Sleep(time.Second * 1)
	}
}

func listenUDP(port string) {
	addr, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

}
