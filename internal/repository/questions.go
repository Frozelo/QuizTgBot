package repository

import (
	"errors"
	"quiz-bot/internal/domain/models"
)

type InMemoryQuestionRepository struct {
	questions []models.Question
}

func NewInMemoryQuestionRepository() *InMemoryQuestionRepository {
	return &InMemoryQuestionRepository{
		questions: []models.Question{
			{
				ID:       1,
				Category: "Go1",
				Question: "Что такое goroutine в Go?",
				Answer:   "Goroutine - это легковесный поток, управляемый Go runtime.",
				Points:   1,
			},
			{
				ID:       2,
				Category: "Go",
				Question: "Что такое канал (channel) в Go?",
				Answer:   "Канал (channel) - это средство для коммуникации между goroutine'ами, обеспечивающее безопасную передачу данных.",
				Points:   2,
			},
			{
				ID:       3,
				Category: "Go",
				Question: "Какие типы данных поддерживаются в Go?",
				Answer:   "Go поддерживает базовые типы данных, такие как int, float64, string, а также сложные типы, такие как struct, array, slice и map.",
				Points:   3,
			},
			{
				ID:       4,
				Category: "Go",
				Question: "Что такое интерфейсы (interfaces) в Go?",
				Answer:   "Интерфейсы (interfaces) в Go - это типы, определяющие поведение через методы, которые должны быть реализованы.",
				Points:   4,
			},
			{
				ID:       5,
				Category: "Go",
				Question: "Как объявить и инициализировать переменную в Go?",
				Answer:   "Переменную в Go можно объявить с помощью ключевого слова var, например: var x int = 10. Также можно использовать короткое объявление: x := 10.",
				Points:   1,
			},
			{
				ID:       6,
				Category: "Go",
				Question: "Что такое пакеты (packages) в Go?",
				Answer:   "Пакеты (packages) в Go - это способ организации кода в независимые модули, что упрощает управление зависимостями и повторное использование кода.",
				Points:   2,
			},
			{
				ID:       7,
				Category: "Go",
				Question: "Что такое defer в Go?",
				Answer:   "defer в Go - это ключевое слово, которое откладывает выполнение функции до момента завершения окружающей функции.",
				Points:   3,
			},
			{
				ID:       8,
				Category: "Go",
				Question: "Что такое panic и recover в Go?",
				Answer:   "panic и recover - это механизмы для обработки ошибок. panic используется для остановки выполнения, а recover позволяет восстановиться после panic.",
				Points:   4,
			},
			{
				ID:       9,
				Category: "Go",
				Question: "Как объявить функцию в Go?",
				Answer:   "Функция в Go объявляется с помощью ключевого слова func, например: func add(a int, b int) int { return a + b }.",
				Points:   1,
			},
			{
				ID:       10,
				Category: "Go",
				Question: "Что такое пустая структура (empty struct) в Go и зачем она используется?",
				Answer:   "Пустая структура (empty struct) - это структура без полей. Она используется для создания значений с нулевым размером, например, в качестве сигнального канала.",
				Points:   2,
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
