package main

import (
	"network_tool/scanners"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new Fyne application and window
	a := app.New()
	mainWindow := a.NewWindow("Scanner")
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.CenterOnScreen()

	// Create channels for communication between goroutines
	tcpConn := make(chan string)
	statusChan := make(chan string)
	done := make(chan bool)

	// Create a label to display TCP connection status
	tcpConnLabel := widget.NewLabel("")

	// Create a waitgroup to keep track of TCP checks
	wg := &sync.WaitGroup{}

	// Create a button to start the TCP scan
	scanButton := widget.NewButton("Scan", func() {
		// Update the label to indicate the scan has started
		tcpConnLabel.SetText("Start.\n")

		// Add a new task to the waitgroup and start a new goroutine to perform the TCP check
		wg.Add(1)
		go func() {
			scanners.TcpCheck(tcpConn, wg, statusChan) // pass the statusChan channel
			done <- true
		}()
	})

	// Create a vertical box to hold the scan button
	sideBar := container.NewVBox(
		scanButton,
	)
	sideBar.Resize(fyne.NewSize(150, 0))

	// Create a scrollable container to hold the TCP connection status label
	tcpConnScroll := container.NewScroll(tcpConnLabel)

	// Create a horizontal split container to display the sidebar and TCP connection status label
	content := container.NewHSplit(
		sideBar,
		tcpConnScroll,
	)
	content.SetOffset(0.2)

	// Set the content of the main window to the horizontal split container
	mainWindow.SetContent(content)

	// Wait for the TCP check to finish and mark the waitgroup as done
	go func() {
		<-done
		wg.Done()
	}()

	// Start a new goroutine to update the TCP connection status label when a message is received on the status channel
	go scanners.UpdateTcpConnLabel(statusChan, tcpConnLabel, done) // pass the statusChan channel

	// Show the main window and start the Fyne event loop
	mainWindow.ShowAndRun()
}
