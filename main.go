// @title Quiz Service API
// @version 1.0
// @description API untuk sistem quiz online
// @host localhost:8080
// @BasePath /
package main

import (
	"log"
	"net/http"

	_ "go-backend-univ/docs" // generated swagger

	httpSwagger "github.com/swaggo/http-swagger"

	"go-backend-univ/db"
	"go-backend-univ/handler"
	"go-backend-univ/repository"
	"go-backend-univ/service"
)

func main() {
	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewQuizRepository(database)
	service := service.NewQuizService(repo)
	handler := handler.NewQuizHandler(service)

	http.HandleFunc("/quiz/start", handler.StartQuiz)
	http.HandleFunc("/quiz/submit", handler.SubmitAnswer)
	http.HandleFunc("/quiz/result", handler.GetResult)
	http.Handle("/swagger/",
		httpSwagger.WrapHandler)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
