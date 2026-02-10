package main

import (
	"fmt"
	"net"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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

func listenUDP(port string, app *tview.Application, peerList *tview.List, input *tview.InputField, finput *tview.InputField) {
	addr, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// ipadd := make(chan string)
	// go sendtcp(ipadd, finput)

	buf := make([]byte, 1024)

	peers := make(map[string]time.Time)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			now := time.Now()
			for ip, t := range peers {
				if now.Sub(t) > 2*time.Second {
					delete(peers, ip)
				}
			}
		}
	}()
	go func() {
		// var ipconn string
		// for {
		// 	fmt.Println("Enter address to connect to :")
		// 	fmt.Scanf("%s", &ipconn)
		// 	if ipconn != "" {
		// 		ipadd <- ipconn
		// 		return
		// 	}
		// }
		input.SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				// text := input.GetText()
				// peerList.AddItem("Selected "+text, "", 0, nil)
				// ipadd <- text
				app.SetFocus(finput)
				// app.Stop()
			}
		})
		finput.SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				aadr := input.GetText()
				path := finput.GetText()
				go sendtcp(aadr, path)
				input.SetText("")
				finput.SetText("")
				app.SetFocus(input)
			}
		})
	}()
	//go func() {
	fmt.Println("Active peers:")
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		peers[string(buf[:n])] = time.Now()

		// for k, _ := range peers {
		// 	fmt.Println(" ", k)
		// }
		app.QueueUpdateDraw(func() {
			peerList.Clear()
			for k := range peers {
				peerList.AddItem(k, "", 0, nil)
			}
		})

	}
	//}()
}
