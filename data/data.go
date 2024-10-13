package data

import (
	"github.com/NicolasBors/world-capitals-quiz/constants"
	"github.com/NicolasBors/world-capitals-quiz/types"
)

var QuestionsData = []types.Question{
	{
		Question:           "What is the capital of France?",
		Options:            [constants.OPTIONS_LENGTH]string{"Lyons", "Marseilles", "Paris"},
		CorrectOptionIndex: 2,
	},
	{
		Question:           "What is the capital of Spain?",
		Options:            [constants.OPTIONS_LENGTH]string{"Barcelona", "Madrid", "Valencia"},
		CorrectOptionIndex: 1,
	},
	{
		Question:           "What is the capital of Swiss?",
		Options:            [constants.OPTIONS_LENGTH]string{"Bern", "Geneva", "Zurich"},
		CorrectOptionIndex: 0,
	},
	{
		Question:           "What is the capital of Morocco?",
		Options:            [constants.OPTIONS_LENGTH]string{"Casablanca", "Marrakech", "Rabat"},
		CorrectOptionIndex: 2,
	},
	{
		Question:           "What is the capital of Brazil?",
		Options:            [constants.OPTIONS_LENGTH]string{"Brasilia", "Rio de Janeiro", "São Paulo"},
		CorrectOptionIndex: 0,
	},
	{
		Question:           "What is the capital of Australia?",
		Options:            [constants.OPTIONS_LENGTH]string{"Canberra", "Melbourne", "Sidney"},
		CorrectOptionIndex: 0,
	},
	{
		Question:           "What is the capital of Canada?",
		Options:            [constants.OPTIONS_LENGTH]string{"Montreal", "Ottawa", "Toronto"},
		CorrectOptionIndex: 2,
	},
	{
		Question:           "What is the capital of the United States?",
		Options:            [constants.OPTIONS_LENGTH]string{"Los Angeles", "New York City", "Washington"},
		CorrectOptionIndex: 2,
	},
	{
		Question:           "What is the capital of Poland?",
		Options:            [constants.OPTIONS_LENGTH]string{"Kraków", "Łódź", "Warsaw"},
		CorrectOptionIndex: 2,
	},
	{
		Question:           "What is the capital of Scotland?",
		Options:            [constants.OPTIONS_LENGTH]string{"Aberdeen", "Edinburgh", "Glasgow"},
		CorrectOptionIndex: 1,
	},
}
