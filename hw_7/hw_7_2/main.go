package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func randomize(randChan chan int, signalMinMaxChan chan bool, resMinMaxChan chan []int, done chan<- bool) {
	for i := 0; i < 10; i++ {
		min := 10
		max := 30
		randomNum := rand.Intn(max-min+1) + min
		randChan <- randomNum
		fmt.Printf("Random num is: %d. ", randomNum)

		signalMinMaxChan <- true
		minMax := <-resMinMaxChan
		fmt.Printf("Smallest number: %d, Largest number: %d\n", minMax[0], minMax[1])

		time.Sleep(time.Second)
	}
	close(randChan)
	done <- true
}

func minMaxFind(randChan chan int, signalMinMaxChan chan bool, resMinMaxChan chan []int) {
	min, max := math.MaxInt, math.MinInt
	for {
		select {
		case num, ok := <-randChan:
			if !ok {
				close(resMinMaxChan)
				return
			}
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		case <-signalMinMaxChan:
			resMinMaxChan <- []int{min, max}
		}
	}
}

func main() {
	randChan := make(chan int)
	signalMinMaxChan := make(chan bool)
	resMinMaxChan := make(chan []int)
	done := make(chan bool)

	go randomize(randChan, signalMinMaxChan, resMinMaxChan, done)
	go minMaxFind(randChan, signalMinMaxChan, resMinMaxChan)

	<-done
}
