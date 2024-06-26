package storage

import "quiz-bot/internal/domain/models"

var Questions = []models.Question{
	{
		ID:       1,
		Category: "Go1",
		Question: "Что такое goroutine в Go?",
		Answers: []models.Answer{
			{
				ID:   1,
				Text: "Легковесный поток (goroutine)",
			},
			{
				ID:   2,
				Text: "Поток (thread)",
			},
			{
				ID:   3,
				Text: "Процесс (process)",
			},
		},
		RightAnswerID: 1,
		Points:        1,
	},
	{
		ID:       2,
		Category: "Go",
		Question: "Что такое канал (channel) в Go?",
		Answers: []models.Answer{
			{
				ID:   1,
				Text: "Легковесный поток (goroutine)",
			},
			{
				ID:   2,
				Text: "Средство общения между каналами",
			},
			{
				ID:   3,
				Text: "Процесс (process)",
			},
		},
		RightAnswerID: 1,
		Points:        2,
	},
	// {
	// 	ID:       3,
	// 	Category: "Go",
	// 	Question: "Какие типы данных поддерживаются в Go?",
	// 	Answers: []mode
	// 	Points:   3,
	// },
	{
		ID:       4,
		Category: "Go",
		Question: "Что такое интерфейсы (interfaces) в Go?",
		Answers: []models.Answer{
			{
				ID:   1,
				Text: "Интерфейс - это набор методов",
			},
			{
				ID:   2,
				Text: "Интерфейс - это нadasdasdasdasdadasd",
			},
		},
		RightAnswerID: 1,
		Points:        4,
	},
	// {
	// 	ID:       5,
	// 	Category: "Go",
	// 	Question: "Как объявить и инициализировать переменную в Go?",
	// 	Answer:   "Переменную в Go можно объявить с помощью ключевого слова var, например: var x int = 10. Также можно использовать короткое объявление: x := 10.",
	// 	Points:   1,
	// },
	// {
	// 	ID:       6,
	// 	Category: "Go",
	// 	Question: "Что такое пакеты (packages) в Go?",
	// 	Answer:   "Пакеты (packages) в Go - это способ организации кода в независимые модули, что упрощает управление зависимостями и повторное использование кода.",
	// 	Points:   2,
	// },
	// {
	// 	ID:       7,
	// 	Category: "Go",
	// 	Question: "Что такое defer в Go?",
	// 	Answer:   "defer в Go - это ключевое слово, которое откладывает выполнение функции до момента завершения окружающей функции.",
	// 	Points:   3,
	// },
	// {
	// 	ID:       8,
	// 	Category: "Go",
	// 	Question: "Что такое panic и recover в Go?",
	// 	Answer:   "panic и recover - это механизмы для обработки ошибок. panic используется для остановки выполнения, а recover позволяет восстановиться после panic.",
	// 	Points:   4,
	// },
	// {
	// 	ID:       9,
	// 	Category: "Go",
	// 	Question: "Как объявить функцию в Go?",
	// 	Answer:   "Функция в Go объявляется с помощью ключевого слова func, например: func add(a int, b int) int { return a + b }.",
	// 	Points:   1,
	// },
	// {
	// 	ID:       10,
	// 	Category: "Go",
	// 	Question: "Что такое пустая структура (empty struct) в Go и зачем она используется?",
	// 	Answer:   "Пустая структура (empty struct) - это структура без полей. Она используется для создания значений с нулевым размером, например, в качестве сигнального канала.",
	// 	Points:   2,
	// },
}
