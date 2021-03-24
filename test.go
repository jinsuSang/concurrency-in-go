package main

import "fmt"

func main() {
	var sum int
	channel := make(chan int)

	go func() {
		channel <- 1
		close(channel)
	}()

	for {
		if item, ok := <-channel; ok {
			sum += item
		} else {
			break
		}
	}
	fmt.Println(sum)
}
