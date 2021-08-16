package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go service1(c1)
	go service2(c2)

	select {
	case res1 := <-c1:
		fmt.Println("Receveived from service 1", res1)
	case res2 := <-c2:
		fmt.Println("Received from service 2", res2)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout, default ")

	}

}

func service1(c chan string) {
	c <- "Response from service 1"
}

func service2(c chan string) {
	c <- "Response from service 2"

}
