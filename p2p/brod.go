package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
func GetBroadcast() (string, error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {

			ipnet, ok := addr.(*net.IPNet)
			if !ok || ipnet.IP.To4() == nil {
				continue
			}

			ip := ipnet.IP.To4()
			mask := ipnet.Mask

			broadcast := make(net.IP, 4)

			for i := 0; i < 4; i++ {
				broadcast[i] = ip[i] | ^mask[i]
			}

			return broadcast.String(), nil
		}
	}

	return "", fmt.Errorf("no network found")
}

type PeerManager struct {
	peers map[string]time.Time
	mu    sync.Mutex
}

func peerUpdater(app *App, pm *PeerManager) {

	for {

		time.Sleep(2 * time.Second)

		now := time.Now()

		pm.mu.Lock()

		list := []string{}

		for ip, t := range pm.peers {

			if now.Sub(t) > 2*time.Second {
				delete(pm.peers, ip)
			} else {
				list = append(list, ip)
			}

		}

		pm.mu.Unlock()

		runtime.EventsEmit(app.ctx, "peer_list", list)
	}
}
func listenUDP(app *App, pm *PeerManager) {

	addr, _ := net.ResolveUDPAddr("udp4", ":8080")
	conn, _ := net.ListenUDP("udp4", addr)

	buf := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		ip := string(buf[:n])

		pm.mu.Lock()
		pm.peers[ip] = time.Now()
		pm.mu.Unlock()
	}
}
