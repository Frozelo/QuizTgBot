package models

type Answer struct {
	ID   int
	Text string
}
type Question struct {
	ID            int
	Category      string
	Question      string
	Answers       []Answer
	RightAnswerID int
	Points        uint
}
