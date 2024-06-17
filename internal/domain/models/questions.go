package models

type Question struct {
	ID       uint
	Category string
	Question string
	Answer   string
	Points   uint
}
