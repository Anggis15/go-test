package service

import (
	"context"
	"go-backend-univ/model"
	"go-backend-univ/repository"
)

type QuizService struct {
	repo *repository.QuizRepository
}

func NewQuizService(repo *repository.QuizRepository) *QuizService {
	return &QuizService{repo: repo}
}

func (s *QuizService) StartQuiz(ctx context.Context, quizID, studentID int) (*model.StartQuizResponse, error) {
	return s.repo.StartQuiz(ctx, quizID, studentID)
}

func (s *QuizService) BulkSubmit(
	ctx context.Context,
	req *model.BulkSubmitRequest,
) (*model.BulkSubmitResponse, error) {
	return s.repo.BulkSubmit(ctx, req)
}

func (s *QuizService) GetResult(ctx context.Context, attemptID int64) (float64, string, error) {
	return s.repo.GetResult(ctx, attemptID)
}
