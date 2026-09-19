package rest

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/first-app/rest/controllers"
	"example.com/first-app/rest/handlers"
	"example.com/first-app/rest/service"
)

func app() {
	connStr := `some string`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	repo := &controllers.PostgresEmployee{DB: db}
	service := &service.EmployeeService{Repo: repo}
	handler := &handlers.EmployeeHandler{Service: *service}
	fmt.Print(handler)
}
