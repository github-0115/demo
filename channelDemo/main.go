package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan int, 1)

	go Producer(queue, 10)
	go Consumer(queue, 10)

	time.Sleep(3 * time.Second)
}

func Producer(queue chan<- int, max int) {
	defer close(queue)
	for i := 0; i < max; i++ {

		if i == max/2 {
			time.Sleep(time.Second * 2)
			continue
		}

		queue <- i
	}
}

func Consumer(queue <-chan int, max int) {

	for i := 0; i < max; i++ {
		select {
		case v := <-queue:
			fmt.Println("consumer:", v)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout")
		}
	}
}
