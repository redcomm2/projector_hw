package entity

type Question struct {
	Question string
	Answers  [4]Answer
	Correct  int
}
