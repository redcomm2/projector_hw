package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomize(randChan chan int) {
	for i := 0; i < 10; i++ {
		min := 10
		max := 30
		randomNum := rand.Intn(max-min+1) + min
		randChan <- randomNum
		fmt.Printf("Random num is: %d. ", randomNum)
		time.Sleep(time.Second)
	}
	close(randChan)
}

func averageCalculate(randChan chan int, avgChan chan float64) {
	var numbers []int
	for num := range randChan {
		numbers = append(numbers, num)
		sum := 0
		for _, v := range numbers {
			sum += v
		}
		avg := float64(sum) / float64(len(numbers))
		avgChan <- avg
	}
	close(avgChan)
}

func print(avgChan chan float64, done chan<- bool) {
	for avg := range avgChan {
		fmt.Printf("Average: %f\n", avg)
	}
	done <- true
}

func main() {
	randChan := make(chan int)
	avgChan := make(chan float64)
	done := make(chan bool)

	go randomize(randChan)
	go averageCalculate(randChan, avgChan)
	go print(avgChan, done)

	<-done
}
