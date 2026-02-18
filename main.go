package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// rules()
	// go brodudp()
	// fmt.Println("Enter 0 for sending files and 1 for recieving files :")
	// var key int
	// fmt.Scanf("%d", &key)
	// if key == 1 {
	// 	go listentcp()
	// } else {

	// 	app := tview.NewApplication()

	// 	peerList := tview.NewList().
	// 		AddItem("Waiting for peers...", "", 0, nil)

	// 	input := tview.NewInputField().
	// 		SetLabel("Select peer: ")
	// 	finput := tview.NewInputField().
	// 		SetLabel("Enter File path: ")
	// 	layout := tview.NewFlex().
	// 		SetDirection(tview.FlexRow).
	// 		AddItem(peerList, 0, 1, false).
	// 		AddItem(input, 3, 1, true).
	// 		AddItem(finput, 3, 1, false)

	// 	app.SetRoot(layout, true)
	// 	go listenUDP(":8080", app, peerList, input, finput)
	// 	if err := app.Run(); err != nil {
	// 		panic(err)
	// 	}

	// }
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")
	mss := widget.NewLabel("Hello ")
	butt := widget.NewButton("Send", func() {
		mss.SetText("Bye")
	})
	mssx := canvas.NewText("tab1", color.NRGBA{R: 255, G: 100, B: 100, A: 255})
	cont := container.NewVBox(mss, mssx, butt)
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", cont),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationTop)
	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
	//select {}
}
