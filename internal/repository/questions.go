package repository

import "quiz-bot/internal/domain/models"

type InMemoryQuestionRepository struct {
	questions []models.Question
}

func NewInMemoryQuestionRepository() *InMemoryQuestionRepository {
	return &InMemoryQuestionRepository{
		questions: []models.Question{
			{
				ID:       1,
				Question: "Что такое goroutine в Go?",
				Answer:   "Goroutine - это легковесный поток, управляемый Go runtime.",
				Points:   1,
			},
		},
	}
}

func (r *InMemoryQuestionRepository) GetAll() []models.Question {
	return r.questions
}

func (r *InMemoryQuestionRepository) GetByID(id uint) (models.Question, bool) {
	for _, q := range r.questions {
		if q.ID == id {
			return q, true
		}
	}
	return models.Question{}, false
}
