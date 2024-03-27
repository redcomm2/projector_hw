package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func randomizer(randChan chan int) {
	for {
		min := 10
		max := 30
		randChan <- rand.Intn(max-min+1) + min
		time.Sleep(1 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randChan := make(chan int)
	numbers := make([]int, 0)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go randomizer(randChan)

	go func() {
		for {
			select {
			case num := <-randChan:
				numbers = append(numbers, num)
				total := 0
				for _, v := range numbers {
					total += v
				}
				average := float64(total) / float64(len(numbers))
				fmt.Printf("Received %d, Average so far: %.2f\n", num, average)
			case <-c:
				fmt.Println("\nProgram interrupted. Exiting.")
				os.Exit(0)
			}
		}
	}()

	select {}
}
