package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"quiz-bot/internal/models"
)

var questions = []models.Question{
	{
		ID: 1, Question: "What is a goroutine in Go?",
		Answer: "A goroutine is a lightweight thread managed by the Go runtime.",
		Points: 1,
	},
	{ID: 2,
		Question: "How do you handle errors in Go?",
		Answer:   "Errors are handled using the error type and the 'if err != nil' pattern.",
		Points:   1,
	},
}

func getQuestion(w http.ResponseWriter, r *http.Request) {
	questionID := r.URL.Query().Get("id")
	if questionID == "" {
		http.Error(w, "Missing question ID", http.StatusBadRequest)
		return
	}

	for _, question := range questions {
		if fmt.Sprintf("%d", question.ID) == questionID {
			json.NewEncoder(w).Encode(question)
			return
		}
	}

	http.Error(w, "Question not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/questions", getQuestion)
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
