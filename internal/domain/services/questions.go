package services

import (
	"math/rand"
	"quiz-bot/internal/domain/models"
	"strconv"
	"time"
)

type ImMemmoryRepository interface {
	GetAll() []models.Question
	GetByID(id int) (models.Question, bool)
	GetAllByCategory(category string) ([]models.Question, error)
	GetCategories() ([]string, error)
}

type QuestionService struct {
	repo ImMemmoryRepository
}

func NewQuestionService(repo ImMemmoryRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) GetRandom() models.Question {
	questions := s.repo.GetAll()
	rand.NewSource(time.Now().UnixMicro())
	return questions[rand.Intn(len(questions))]
}

func (s *QuestionService) GetRandomByCategory(category string) (models.Question, error) {
	question, err := s.repo.GetAllByCategory(category)
	if err != nil {
		return models.Question{}, err
	}
	rand.NewSource(time.Now().UnixMicro())
	return question[rand.Intn(len(question))], nil

}

func (s *QuestionService) GetCategories() ([]string, error) {
	return s.repo.GetCategories()
}

func (s *QuestionService) GetRightAnswerID(question *models.Question) int {
	return question.RightAnswerID

}

func (s *QuestionService) GetRightAnswerText(question *models.Question) string {
	for _, answer := range question.Answers {
		if answer.ID == question.RightAnswerID {
			return answer.Text
		}
	}
	return ""
}

func (s *QuestionService) CheckUserAnswer(question *models.Question, answer string) (bool, error) {
	userAnswer, err := strconv.Atoi(answer)
	if err != nil {
		return false, err
	}

	return s.GetRightAnswerID(question) == userAnswer, nil
}
