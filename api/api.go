package main

import (
	"fmt"
	"net/http"

	"github.com/NicolasBors/world-capitals-quiz/data"
	"github.com/gin-gonic/gin"
)

func getQuestions(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, data.QuestionsData)
}

func postAnswers(context *gin.Context) {
	var answers []int
	if err := context.BindJSON(&answers); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request: %s", err),
		})
		return
	}

	correctAnswers := 0
	for index, answer := range answers {
		if data.QuestionsData[index].CorrectOptionIndex == answerToOptionIndex(answer) {
			correctAnswers++
		}
	}

	score := Score{Value: correctAnswers}
	response := fmt.Sprintf("Your score is %d/%d. You did better than %d%% of other players!", score.Value, len(data.QuestionsData), int(score.BetterThanPercent()))
	score.Record()

	context.IndentedJSON(http.StatusOK, response)
}
