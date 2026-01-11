package model

type Question struct {
	ID      int    `json:"question_id"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	Options string `json:"options_json"`
}

type StartQuizResponse struct {
	AttemptID int64      `json:"attempt_id"`
	Questions []Question `json:"questions"`
}
