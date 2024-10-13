package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NicolasBors/world-capitals-quiz/constants"
	"github.com/NicolasBors/world-capitals-quiz/data"
	"github.com/NicolasBors/world-capitals-quiz/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET(constants.PATH_GET_QUESTIONS, getQuestions)
	return router
}

func TestGetQuestions(t *testing.T) {
	var (
		recorder = httptest.NewRecorder()
		req, _   = http.NewRequest(http.MethodGet, constants.PATH_GET_QUESTIONS, nil)
		router   = setupRouter()
	)

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var questions []types.Question
	err := json.Unmarshal(recorder.Body.Bytes(), &questions)
	assert.Nil(t, err)

	assert.Equal(t, len(data.QuestionsData), len(questions), "Expected the number of questions to match")
}
