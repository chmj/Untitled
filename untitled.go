package main

import (
	"fmt"
	"time"
)

/*
 The kicking order after the referee kicks-off the game is random (0_0).
 More like "who gets the ball first" from the referee.
*/

func main() {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Println(<-ball, "kicked the ball.")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("John")
	go kickBall("Alice")
	go kickBall("Bob")
	go kickBall("Emily")
	ball <- "referee" // kick off
	var c chan bool   // nil
	<-c               // blocking here forever
}
