package main

import (
	"fmt"
	"projector_hw/hw_8/entity"
	"sync"
	"time"
)

func generateRound(startChan chan int, rounds []entity.Round, wg *sync.WaitGroup) {
	<-startChan
	for _, round := range rounds {
		round.Process(wg)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	startChan := make(chan int)
	questions := entity.GetQuestions()
	var wg sync.WaitGroup

	rounds := make([]entity.Round, 7)
	wg.Add(len(rounds))
	for i := 0; i < 7; i++ {
		question, exists := questions[i+1]
		if exists {
			rounds[i] = entity.Round{ID: i + 1, Question: question}
		} else {
			fmt.Printf("Question for round %d not found.\n", i+1)
		}
	}

	go generateRound(startChan, rounds, &wg)
	startChan <- 1
	wg.Wait()
}
