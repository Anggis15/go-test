package model

type StartQuizRequest struct {
	QuizID int `json:"quiz_id"`
}

type SubmitAnswerRequest struct {
	AttemptID  int64  `json:"attempt_id"`
	QuestionID int    `json:"question_id"`
	AnswerText string `json:"answer_text"`
}

type QuizResult struct {
	FinalScore float64 `json:"final_score"`
	Status     string  `json:"status"`
}

type BulkSubmitRequest struct {
	AttemptID int64 `json:"attempt_id"`
	Answers   []struct {
		QuestionID int     `json:"question_id"`
		AnswerText *string `json:"answer_text,omitempty"`
		FileURL    *string `json:"file_url,omitempty"`
	} `json:"answers"`
}

type BulkSubmitResponse struct {
	Status     string   `json:"status"`
	FinalScore *float64 `json:"final_score"`
}
