package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func tcpCheck(tcpConn chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1024; i++ {
		tcpConnStr := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", tcpConnStr)
		if err != nil {
			// Send the closed port message back to the main goroutine via the channel
			tcpConn <- fmt.Sprintf("port "+"%d closed\n", i)
			continue
		}
		conn.Close()

		// Send the open port message back to the main goroutine via the channel
		tcpConn <- fmt.Sprintf("port "+"%d open\n", i)

		// Add a delay to slow down the function
		time.Sleep(10 * time.Millisecond)
	}
}

func updateTcpConnLabel(tcpConn chan string, tcpConnLabel *widget.Label, done chan bool) {
	tcpConnLabel.SetText("")
	i := 0
	for msg := range tcpConn {
		parts := strings.Split(msg, " ")
		if len(parts) != 2 {
			continue
		}
		port, status := parts[0], parts[1]
		tcpConnLabel.SetText(tcpConnLabel.Text + fmt.Sprintf("%s %s\n", port, status))

		i++
	}
	tcpConnLabel.SetText(tcpConnLabel.Text + "Done.\n")

	done <- true
}

func main() {
	a := app.New()
	mainWindow := a.NewWindow("Scanner")
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.CenterOnScreen()

	tcpConn := make(chan string)
	done := make(chan bool)
	tcpConnLabel := widget.NewLabel("")
	wg := &sync.WaitGroup{}

	scanButton := widget.NewButton("Scan", func() {
		tcpConnLabel.SetText("Start.\n")

		// Reset the WaitGroup before starting the scan
		wg.Add(1)

		go func() {
			tcpCheck(tcpConn, wg) // add the WaitGroup pointer as the second argument
			done <- true
		}()
	})

	sideBar := container.NewVBox(
		scanButton,
	)
	sideBar.Resize(fyne.NewSize(150, 0))

	tcpConnScroll := container.NewScroll(tcpConnLabel)

	content := container.NewHSplit(
		sideBar,
		tcpConnScroll,
	)
	content.SetOffset(0.2)

	mainWindow.SetContent(content)

	// Wait for tcpCheck to finish before closing the WaitGroup
	go func() {
		<-done
		wg.Done()
	}()

	go updateTcpConnLabel(tcpConn, tcpConnLabel, done)

	mainWindow.ShowAndRun()
}
