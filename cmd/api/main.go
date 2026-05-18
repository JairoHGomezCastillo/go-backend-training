package main

import (
	"database/sql"
	"go-backend-training/internal/handlers"
	"go-backend-training/internal/stack"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/stacksdbapi?parseTime=true")
	if err != nil {
		return err
	}
	repo := stack.NewStackRepository(db)
	stackUseCase := stack.NewStackUseCase(repo)
	stackHandler := handlers.NewStackHandler(stackUseCase)

	http.HandleFunc("POST /application/{applicationName}/stacks", stackHandler.CreateStackHandler)
	http.HandleFunc("GET /application/{applicationName}/stacks/{stackName}", stackHandler.SearchStackByNameHandler)

	return http.ListenAndServe("localhost:8080", nil)
}
