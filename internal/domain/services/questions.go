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

func (s *QuestionService) CheckAnswer(question models.Question, answer string) (bool, error) {
	// expectedWords := strings.Fields(strings.ToLower(question.Answer))
	// userWords := strings.Fields(strings.ToLower(answer))

	// words := make(map[string]bool)
	// for _, word := range expectedWords {
	// 	words[word] = true
	// }
	// for _, word := range userWords {
	// 	if _, ok := words[word]; ok {
	// 		return true
	// 	}
	// }
	// return false
	userAnswer, err := strconv.Atoi(answer) // Преобразуем ответ пользователя в число
	if err != nil {
		return false, err
	}

	return question.RightAnswerID == userAnswer, nil
}
