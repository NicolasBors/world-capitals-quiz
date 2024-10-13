package types

type Question struct {
	CorrectOptionIndex int
	Options            [3]string `json:"options"`
	Question           string    `json:"question"`
}
