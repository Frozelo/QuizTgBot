package services

import (
	"math/rand"
	"quiz-bot/internal/domain/models"
	"strings"
	"time"
)

type ImMemmoryRepository interface {
	GetAll() []models.Question
	GetByID(id uint) (models.Question, bool)
}

type QuestionService struct {
	repo ImMemmoryRepository
}

func NewQuestionService(repo ImMemmoryRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) GetRandomQuestion() models.Question {
	questions := s.repo.GetAll()
	rand.NewSource(time.Now().UnixMicro())
	return questions[rand.Intn(len(questions))]
}

func (s *QuestionService) CheckAnswer(question models.Question, answer string) bool {
	expectedAnswer := question.Answer
	expectedWords := strings.Fields(expectedAnswer)
	userWords := strings.Fields(answer)

	for _, word := range expectedWords {
		for _, userWord := range userWords {
			if strings.Contains(strings.ToLower(userWord), strings.ToLower(word)) {
				return true
			}
		}
	}

	return false
}
