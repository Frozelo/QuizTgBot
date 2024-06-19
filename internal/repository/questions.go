package repository

import (
	"errors"
	"quiz-bot/internal/domain/models"
	"quiz-bot/internal/storage"
)

type InMemoryQuestionRepository struct {
	questions []models.Question
}

func NewInMemoryQuestionRepository() *InMemoryQuestionRepository {
	return &InMemoryQuestionRepository{
		questions: storage.Questions,
	}
}

func (r *InMemoryQuestionRepository) GetAll() []models.Question {
	return r.questions
}

func (r *InMemoryQuestionRepository) GetByID(id int) (models.Question, bool) {
	for _, q := range r.questions {
		if q.ID == id {
			return q, true
		}
	}
	return models.Question{}, false
}

func (r *InMemoryQuestionRepository) GetAllByCategory(category string) ([]models.Question, error) {
	questions := []models.Question{}
	for _, q := range r.questions {
		if q.Category == category {
			questions = append(questions, q)
		}
	}
	if len(questions) == 0 {
		return []models.Question{}, errors.New("No Question found for category " + category)
	}
	return questions, nil
}

func (r *InMemoryQuestionRepository) GetByCategory(category string) (models.Question, error) {
	for _, q := range r.questions {
		if q.Category == category {
			return q, nil
		}
	}
	return models.Question{}, nil
}

func (r *InMemoryQuestionRepository) GetCategories() ([]string, error) {
	categories := make([]string, 0)
	seenCategories := make(map[string]struct{}, 0)

	for _, q := range r.questions {
		if _, ok := seenCategories[q.Category]; !ok {
			categories = append(categories, q.Category)
			seenCategories[q.Category] = struct{}{}
		}
	}
	if len(categories) == 0 {
		return []string{}, errors.New("no categories found")
	}
	return categories, nil
}
