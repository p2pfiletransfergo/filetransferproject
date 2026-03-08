package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	pm := &PeerManager{
		peers: make(map[string]time.Time),
	}
	go listenUDP(a, pm)
	go peerUpdater(a, pm)
	go rules()
}

func (a *App) Sendtcp(addr string, path string) {
	time.Sleep(2 * time.Second)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
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

	var resp uint8
	binary.Read(conn, binary.BigEndian, &resp)

	if resp == 0 {
		fmt.Println("Receiver rejected file")
		return
	}
	io.Copy(conn, file)

}

func (a *App) SelectFile() (string, error) {

	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
	})

	return path, err
}

func Handleconn(conn net.Conn, a *App) {
	defer conn.Close()
	PreventSleep()
	defer AllowSleep()
	var nameLen uint32
	binary.Read(conn, binary.BigEndian, &nameLen)

	nameBuf := make([]byte, nameLen)
	io.ReadFull(conn, nameBuf)
	filename := string(nameBuf)

	var fileSize int64
	binary.Read(conn, binary.BigEndian, &fileSize)
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	allowed := a.AskPermission(id, filename, fileSize)

	if !allowed {
		binary.Write(conn, binary.BigEndian, uint8(0))
		conn.Close()
		return
	}
	binary.Write(conn, binary.BigEndian, uint8(1))
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

	buffer := make([]byte, 4096)
	var received int64 = 0

	for received < fileSize {

		n, err := conn.Read(buffer)
		if err != nil {
			break
		}

		file.Write(buffer[:n])
		received += int64(n)

		progress := int(float64(received) / float64(fileSize) * 100)

		runtime.EventsEmit(a.ctx, "recv_progress", map[string]interface{}{
			"id":       id,
			"progress": progress,
			"name":     filename,
		})
	}

	fmt.Println("Received:", filename)
}

var pendingPermissions = map[string]chan bool{}
var permMutex sync.Mutex

func (a *App) AllowFile(id string) {

	permMutex.Lock()
	ch, ok := pendingPermissions[id]
	permMutex.Unlock()

	if ok {
		ch <- true
	}
}
func (a *App) RejectFile(id string) {

	permMutex.Lock()
	ch, ok := pendingPermissions[id]
	permMutex.Unlock()

	if ok {
		ch <- false
	}
}
func (a *App) AskPermission(id string, name string, size int64) bool {

	ch := make(chan bool)

	permMutex.Lock()
	pendingPermissions[id] = ch
	permMutex.Unlock()

	runtime.EventsEmit(a.ctx, "file_request", map[string]interface{}{
		"id":   id,
		"name": name,
		"size": size,
	})

	resp := <-ch

	permMutex.Lock()
	delete(pendingPermissions, id)
	permMutex.Unlock()

	return resp
}
func (a *App) Listentcp() {
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
		go Handleconn(conn, a)
	}
}

func (a *App) Brodudp() {
	badd, err := GetBroadcast()
	if err != nil {
		panic(err)
	}
	raddr, _ := net.ResolveUDPAddr("udp4", badd+":8080")

	conn, _ := net.DialUDP("udp4", nil, raddr)
	conn.SetWriteBuffer(1024)

	for {
		p := getIP()
		conn.Write([]byte(p.IP.String() + ":5050"))
		time.Sleep(time.Second * 1)
	}
}
