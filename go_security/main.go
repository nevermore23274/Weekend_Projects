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

/*
// tcpCheck scans ports 1 through 1024 of "scanme.nmap.org".
// It sends a string with the format "{port number} open/closed" to statusChan for each port scanned.
// If an error occurs while attempting to connect to a port, no message is sent to statusChan.
// wg is used to coordinate goroutines
*/
func tcpCheck(tcpConn chan string, wg *sync.WaitGroup, statusChan chan string) {
	defer wg.Done()
	for i := 1; i <= 1024; i++ {
		tcpConnStr := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", tcpConnStr)
		if err != nil {
			// Send nothing to tcpConn channel if the port is closed
			statusChan <- fmt.Sprintf("%d closed", i)
			continue
		}
		conn.Close()

		// Send the open port message back to the main goroutine via the statusChan
		statusChan <- fmt.Sprintf("%d open", i)

		// Add a delay to slow down the function
		time.Sleep(10 * time.Millisecond)
	}
}

func updateTcpConnLabel(statusChan chan string, tcpConnLabel *widget.Label, done chan bool) {
	tcpConnLabel.SetText("Start.\n") // Set initial label text to indicate that the scan has started
	for msg := range statusChan {    // Loop over messages in the channel until it's closed
		parts := strings.Split(msg, " ") // Split the message into two parts: port number and status
		if len(parts) != 2 {             // If the message doesn't have two parts, skip it
			continue
		}
		port, status := parts[0], parts[1]                                             // Get the port number and status from the message
		tcpConnLabel.SetText(tcpConnLabel.Text + fmt.Sprintf("%s %s\n", port, status)) // Update the label text with the new port and status
	}
	tcpConnLabel.SetText(tcpConnLabel.Text + "Done.\n") // When the channel is closed, indicate that the scan is complete
	done <- true                                        // Signal that the function has finished by sending a value to the done channel
}

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
			tcpCheck(tcpConn, wg, statusChan) // pass the statusChan channel
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
	go updateTcpConnLabel(statusChan, tcpConnLabel, done) // pass the statusChan channel

	// Show the main window and start the Fyne event loop
	mainWindow.ShowAndRun()
}
