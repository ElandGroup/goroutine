package main

import (
	"fmt"
	"time"
)

func Producer(queue chan<- int) {

	for i := 0; i < 10; i++ {

		queue <- i

	}

}

func Consumer(queue <-chan int) {

	for true {

		v := <-queue

		fmt.Println("receive:", v)

	}

}

func ProducerAndConsumer() {

	queue := make(chan int, 1)

	go Consumer(queue)
	go Producer(queue)

	time.Sleep(1e9) //Waiting for Producer and Consumer to complete
	// var input string
	// fmt.Scanln(&input)
}
