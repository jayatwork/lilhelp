//go to town on displaying time mangt

// Golang program to illustrate the usage of
// time.NewTicker() function

// Including main package
package timer

// Importing fmt and time
import (
	"fmt"
	"time"
)

// Calling main
func Start() {

	// Calling NewTicker method
	d := time.NewTicker(2 * time.Second)

	// Creating channel using make
	// keyword
	mychannel := make(chan bool)

	// Calling Sleep() methpod in go
	// function
	go func() {
		time.Sleep(7 * time.Second)

		// Setting the value of channel
		mychannel <- true
	}()

	// Using for loop
	for {

		// Select statement
		select {

		// Case statement
		case <-mychannel:
			fmt.Println("Completed!")
			return

		// Case to print current time
		case tm := <-d.C:
			fmt.Println("The Current time is: ", tm)
		}
	}
}
