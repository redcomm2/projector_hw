package entity

import (
	"fmt"
	"sync"
)

type Round struct {
	ID       int
	Question Question
}

func (r *Round) Process(wg *sync.WaitGroup) {
	fmt.Println(r.ID)
	wg.Done()
}

func GetQuestions() map[int]Question {
	return map[int]Question{
		1: {
			Question: "What is the capital of France?",
			Answers: [4]Answer{
				{Answer: "Paris"},
				{Answer: "Berlin"},
				{Answer: "Madrid"},
				{Answer: "Lisbon"},
			},
			Correct: 0,
		},
		2: {
			Question: "Which element has the chemical symbol 'O'?",
			Answers: [4]Answer{
				{Answer: "Gold"},
				{Answer: "Oxygen"},
				{Answer: "Silver"},
				{Answer: "Helium"},
			},
			Correct: 1,
		},
		3: {
			Question: "Who wrote 'Hamlet'?",
			Answers: [4]Answer{
				{Answer: "Charles Dickens"},
				{Answer: "William Shakespeare"},
				{Answer: "Leo Tolstoy"},
				{Answer: "Mark Twain"},
			},
			Correct: 1,
		},
		4: {
			Question: "What is the largest planet in our solar system?",
			Answers: [4]Answer{
				{Answer: "Earth"},
				{Answer: "Jupiter"},
				{Answer: "Mars"},
				{Answer: "Neptune"},
			},
			Correct: 1,
		},
		5: {
			Question: "Which of these animals is a mammal?",
			Answers: [4]Answer{
				{Answer: "Shark"},
				{Answer: "Dolphin"},
				{Answer: "Crocodile"},
				{Answer: "Octopus"},
			},
			Correct: 1,
		},
		6: {
			Question: "What is the hardest natural substance on Earth?",
			Answers: [4]Answer{
				{Answer: "Gold"},
				{Answer: "Iron"},
				{Answer: "Diamond"},
				{Answer: "Quartz"},
			},
			Correct: 2,
		},
		7: {
			Question: "In which country is the Great Barrier Reef located?",
			Answers: [4]Answer{
				{Answer: "Australia"},
				{Answer: "Brazil"},
				{Answer: "India"},
				{Answer: "South Africa"},
			},
			Correct: 0,
		},
	}

}
