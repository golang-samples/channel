package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits	int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("ping", table)

	table <- new(Ball) // Game on; toss the ball
	time.Sleep(time.Second)
	<-table // Game over; grab the ball
}
