#Soal Tes: Backend Developer

##Task 1
Create Table Student

```bash
CREATE TABLE students (
    student_id INT NOT NULL,
    full_name VARCHAR(100) NOT NULL
);
```

Create Table Questions

```bash
CREATE TABLE questions (
    question_id INT NOT NULL,
    question_type VARCHAR(20) NOT NULL, -- MCQ | ESSAY | FILE
    question_text NVARCHAR(MAX) NOT NULL,
    correct_answer VARCHAR(50) NULL,
    max_score INT NOT NULL
);
```

Create Table Quiz Question

```bash
CREATE TABLE quiz_questions (
    quiz_question_id INT IDENTITY(1,1),
    quiz_id INT NOT NULL,
    question_id INT NOT NULL,
    question_order INT NOT NULL,
    weight DECIMAL(5,2) NOT NULL,
    created_at DATETIME2 NOT NULL DEFAULT SYSDATETIME()
);

```
Create Table Quizzes

```bash
CREATE TABLE quizzes (
    quiz_id INT NOT NULL,
    course_id INT NOT NULL,
    title VARCHAR(200) NOT NULL,
    duration_minutes INT NOT NULL,
    retake_limit INT NOT NULL,
    passing_score INT NOT NULL
);

```
Create Table Quiz Attempts

```bash
CREATE TABLE quiz_attempts (
    attempt_id INT IDENTITY(1,1),
    quiz_id INT NOT NULL,
    student_id INT NOT NULL,
    start_time DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
    end_time DATETIME2 NULL,
    final_score DECIMAL(5,2) NULL,
    status VARCHAR(30) NOT NULL,
    has_manual_question BIT NOT NULL
);

```
Create Table Answeer

```bash
CREATE TABLE answers (
    answer_id INT IDENTITY(1,1),
    attempt_id INT NOT NULL,
    question_id INT NOT NULL,
    answer_text NVARCHAR(MAX) NULL,
    file_url VARCHAR(255) NULL,
    score DECIMAL(5,2) NULL,
    is_correct BIT NULL,
    is_auto_scored BIT NOT NULL
);

```

##Task 2

Ubah Configurasi pada package db pada function newDB()

Buka Menggunakan URL:

```bash
http://localhost:8080/swagger/index.html#/Quiz/post_quiz_submit
```

##Task 3

```python

┌───────────────────────┐
│      START QUIZ       │
│ (Create Quiz Attempt) │
└───────────┬───────────┘
            │
            ▼
┌──────────────────────────────┐
│ Student Submit Answers (Bulk)│
│ - MCQ                        │
│ - Essay                      │
│ - File Upload                │
└───────────┬──────────────────┘
            │
            ▼
┌──────────────────────────────┐
│     Auto Grading MCQ         │
│ - Compare with correct_answer│
│ - Set score & is_correct     │
└───────────┬──────────────────┘
            │
            ▼
┌──────────────────────────────┐
│ Apakah ada soal ESSAY / FILE │
│ yang BELUM dinilai dosen?    │
└───────────┬───────────┬──────┘
            │YES        │NO
            ▼           ▼
┌──────────────────┐   ┌──────────────────────────┐
│ Status Attempt   │   │ Hitung Nilai Akhir       │
│ = WAITING_       │   │ - Sum(score * weight)    │
│   ASSESSMENT     │   └───────────┬──────────────┘
└──────────┬───────┘               │
           │                       ▼
           │            ┌─────────────────────────┐
           │            │ Apakah nilai ≥ passing? │
           │            └───────────┬───────────┬─┘
           │                        │YES        │NO
           ▼                        ▼           ▼
┌────────────────────┐     ┌────────────────┐ ┌────────────────┐
│ Menunggu penilaian │     │ Status = PASSED│ │ Status = FAILED│
│ dosen              │     └────────────────┘ └────────────────┘
└────────────────────┘
```


```bash
SELECT
    s.student_id,
    s.full_name,
    qz.quiz_id,
    qz.title AS quiz_title,
    qa.final_score,
    qa.status AS quiz_status,
    CASE
        WHEN qa.status = 'PASSED' THEN 'LULUS'
        WHEN qa.status = 'FAILED' THEN 'TIDAK LULUS'
        WHEN qa.status = 'WAITING_ASSESSMENT' THEN 'MENUNGGU PENILAIAN'
        ELSE qa.status
    END AS status_kelulusan
FROM quizzes qz
JOIN quiz_attempts qa
    ON qa.quiz_id = qz.quiz_id
JOIN students s
    ON s.student_id = qa.student_id
WHERE qz.course_id = @course_id
ORDER BY
    s.full_name,
    qz.title;
```