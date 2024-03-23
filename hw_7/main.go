package main

import "fmt"

func sendNumbers(ch chan int, done chan bool) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
	done <- true
}

func print(ch chan int, done chan bool) {
	fmt.Println(<-ch)
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go sendNumbers(ch, done)

	go print(ch, done)

	<-done
}
