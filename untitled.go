package main

import (
	"fmt"
	"time"
)

/*
Some Channel Use Examples

Now that you've read the above section, let's view some examples which use channels to enhance your understanding.
A simple request/response example. The two goroutines in this example talk to each other through an unbuffered channel.
*/

func main() {
	c := make(chan int) // an unbuffered channel
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// <-ch    // fails to compile
		// Send the value and block until the result is received.
		ch <- x * x // 9 is sent
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		// Block until 9 is received.
		n := <-ch
		fmt.Println(n) // 9
		// ch <- 123   // fails to compile
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)
	// Block here until a value is received by
	// the channel "done".
	<-done
	fmt.Println("bye")
}
