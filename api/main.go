package main

import (
	"github.com/NicolasBors/world-capitals-quiz/constants"
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()
	router.GET(constants.PATH_GET_QUESTIONS, getQuestions)
	router.POST(constants.PATH_POST_ANSWERS, postAnswers)
	router.Run(":" + constants.SERVER_PORT)
}
