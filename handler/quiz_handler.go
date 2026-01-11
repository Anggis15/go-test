package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-backend-univ/halper"
	"go-backend-univ/model"
	"go-backend-univ/service"
)

type QuizHandler struct {
	service *service.QuizService
}

func NewQuizHandler(service *service.QuizService) *QuizHandler {
	return &QuizHandler{service: service}
}

// StartQuiz godoc
// @Summary Start quiz
// @Description Memulai quiz dan mengembalikan soal
// @Tags Quiz
// @Accept json
// @Produce json
// @Param body body model.StartQuizRequest true "Start Quiz"
// @Router /quiz/start [post]
func (h *QuizHandler) StartQuiz(w http.ResponseWriter, r *http.Request) {
	var req model.StartQuizRequest
	json.NewDecoder(r.Body).Decode(&req)

	studentID := 1 // from JWT

	resp, err := h.service.StartQuiz(r.Context(), req.QuizID, studentID)
	if err != nil {
		halper.WriteError(w, 400, err.Error())
		return
	}

	json.NewEncoder(w).Encode(resp)
}

// SubmitAnswer godoc
// @Summary Submit answer
// @Description Autosave jawaban per soal
// @Tags Quiz
// @Accept json
// @Produce json
// @Param body body model.BulkSubmitRequest true "Submit Answer"
// @Success 200
// @Router /quiz/submit [post]
func (h *QuizHandler) SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	var req model.BulkSubmitRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp, err := h.service.BulkSubmit(r.Context(), &req)
	if err != nil {
		halper.WriteError(w, 400, err.Error())
		return
	}

	json.NewEncoder(w).Encode(resp)
}

// GetResult godoc
// @Summary Get quiz result
// @Description Melihat hasil akhir quiz
// @Tags Quiz
// @Produce json
// @Param attempt_id query int true "Attempt ID"
// @Router /quiz/result [get]
func (h *QuizHandler) GetResult(w http.ResponseWriter, r *http.Request) {
	attemptID, _ := strconv.ParseInt(r.URL.Query().Get("attempt_id"), 10, 64)

	score, status, err := h.service.GetResult(r.Context(), attemptID)
	if err != nil {
		halper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(model.QuizResult{
		FinalScore: score,
		Status:     status,
	})
}
