package controllers

import (
	"database/sql"
	"time"

	"example.com/first-app/rest/models"
)

type PostgresEmployee struct {
	DB *sql.DB
}

func (r *PostgresEmployee) GetAll() ([]models.Employee, error) {
	rows, err := r.DB.Query("select id,first_name,last_name,created_at from employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		err := rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func (r *PostgresEmployee) CreateEmployee(e models.Employee) (models.Employee, error) {
	query := `insert into employees(first_name,last_name,created_at) values($1,$2,$3) returning id`
	err := r.DB.QueryRow(query, &e.FirstName, &e.LastName, time.Now()).Scan(&e.ID)
	if err != nil {
		return e, err
	}
	e.CreatedAt = time.Now()
	return e, nil
}
