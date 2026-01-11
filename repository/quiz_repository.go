package repository

import (
	"context"
	"database/sql"
	"go-backend-univ/model"
)

type QuizRepository struct {
	db *sql.DB
}

func NewQuizRepository(db *sql.DB) *QuizRepository {
	return &QuizRepository{db: db}
}

func (r *QuizRepository) StartQuiz(ctx context.Context, quizID, studentID int) (*model.StartQuizResponse, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"EXEC sp_StartQuiz @StudentId=@p1, @QuizId=@p2",
		studentID,
		quizID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []model.Question

	// Result set 1: questions
	for rows.Next() {
		var q model.Question
		rows.Scan(&q.ID, &q.Type, &q.Text, &q.Options)
		questions = append(questions, q)
	}

	// Move to result set 2
	rows.NextResultSet()

	var attemptID int64
	if rows.Next() {
		rows.Scan(&attemptID)
	}

	return &model.StartQuizResponse{
		AttemptID: attemptID,
		Questions: questions,
	}, nil
}

func (r *QuizRepository) BulkSubmit(ctx context.Context, req *model.BulkSubmitRequest) (*model.BulkSubmitResponse, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, a := range req.Answers {
		_, err := tx.ExecContext(ctx, `
			MERGE answers AS t
			USING (SELECT @p1 attempt_id, @p2 question_id) s
			ON t.attempt_id = s.attempt_id AND t.question_id = s.question_id
			WHEN MATCHED THEN
			  UPDATE SET answer_text = @p3, file_url = @p4
			WHEN NOT MATCHED THEN
			  INSERT (attempt_id, question_id, answer_text, file_url)
			  VALUES (@p1, @p2, @p3, @p4);
		`,
			req.AttemptID,
			a.QuestionID,
			a.AnswerText,
			a.FileURL,
		)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Call grading SP
	var resp model.BulkSubmitResponse
	err = tx.QueryRowContext(ctx,
		"EXEC sp_BulkSubmitQuiz @AttemptId=@p1",
		req.AttemptID,
	).Scan(&resp.Status, &resp.FinalScore)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &resp, nil
}

func (r *QuizRepository) GetResult(ctx context.Context, attemptID int64) (float64, string, error) {
	var score sql.NullFloat64
	var status string

	err := r.db.QueryRowContext(ctx, `
		SELECT final_score, status
		FROM quiz_attempts
		WHERE attempt_id = @p1
	`, attemptID).Scan(&score, &status)

	return score.Float64, status, err
}
