// Go's select lets you wait on multiple channel operations. Combining
// goroutines an channels with select is a powerful feature of Go.
package main

import "time"
import "fmt"

func main() {

	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time, to simulate
	// e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// We'll use select to await both of these values simultaneously,
		// printing each one as it arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)

		}
	}
}
