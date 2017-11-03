// We often want to execute Go code at some point in the future.
// Go's built-in timer feature make this task easy.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers represent a single event in the future, You tell the timer how
	// long you want to wait, and it provides a channel that will be notified
	// at that time. This timer will wait 2 seconds
	timer1 := time.NewTimer(time.Second * 2)

	// The <-timer.C blocks on the timer's channel c until it sends a value
	// indicating that the timer expired.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel
	// the timer before it expires. Here's an example of that
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
}
