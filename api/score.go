package main

import (
	"github.com/NicolasBors/world-capitals-quiz/data"
)

var recordedScores = []int{}

type Score struct {
	Value int
}

func (s *Score) BetterThanPercent() float64 {
	if len(recordedScores) == 0 {
		return 0
	}

	betterCount := 0
	for _, score := range recordedScores {
		if s.Value > score {
			betterCount++
		}
	}

	return float64(betterCount) / float64(len(recordedScores)) * 100
}

func (s *Score) Ratio() float64 {
	return float64(s.Value) / float64(len(data.QuestionsData))
}

func (s *Score) Record() {
	recordedScores = append(recordedScores, s.Value)
}
