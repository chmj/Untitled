package main

import (
	"log"
	"math/rand"
	"time"
)

/*
 Buffered channels can be used as counting semaphores. Counting semaphores can be viewed as multi-owner locks.
 If the capacity of a channel is N, then it can be viewed as a lock which can have most N owners at any time.
 Binary semaphores (mutexes) are special counting semaphores, each of binary semaphores can have at most one owner at
 any time.

 Counting semaphores are often used to enforce a maximum number of concurrent requests.
 Like using channels as mutexes, there are also two manners to acquire one piece of ownership of a channel semaphore.

    Acquire ownership through a send, release through a receive.
    Acquire ownership through a receive, release through a send.

 An example of acquiring ownership through receiving values from a channel.
*/

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("customer#", c, " enters the bar")
	seat := <-bar // need a seat to drink
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- customer#", c, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// the bar has 10 seats.
	bar24x7 := make(Bar, 10)
	// Place seats in an bar.
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// None of the sends will block.
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}

	// sleeping != blocking
	for {
		time.Sleep(time.Second)
	}
}
