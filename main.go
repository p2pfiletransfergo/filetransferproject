package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {

	go brodudp()
	fmt.Println("Enter 0 for sending files and 1 for recieving files :")
	var key int
	fmt.Scanf("%d", &key)
	if key == 1 {
		go listentcp()
	} else {

		app := tview.NewApplication()

		peerList := tview.NewList().
			AddItem("Waiting for peers...", "", 0, nil)

		input := tview.NewInputField().
			SetLabel("Select peer: ")
		finput := tview.NewInputField().
			SetLabel("Enter File path: ")
		layout := tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(peerList, 0, 1, false).
			AddItem(input, 3, 1, true).
			AddItem(finput, 3, 1, false)

		app.SetRoot(layout, true)
		go listenUDP(":8080", app, peerList, input, finput)
		if err := app.Run(); err != nil {
			panic(err)
		}

	}

	select {}
}
