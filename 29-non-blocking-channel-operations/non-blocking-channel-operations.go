/* Basic sends and receives on channels are blocking. However, we can use
select with a default clause to implement non-blocking sends, receives,
and even non-blocking multi-way selects. */
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	/* Here's a non-blocking receive. If value is availabel on messages then
	select will take the <-messages case with that value. If not it will
	immediately take the default case. */
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "h1"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent", msg)
	}

	/* We can multiple cases above the default clause to implement a mult-way
	non-blocking select. Here we attempt non-blocking receives on both messages
	and signals. */
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
