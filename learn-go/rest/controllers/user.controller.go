package controllers

import (
	"database/sql"
	"time"

	"example.com/first-app/rest/models"
)

type PostgresUser struct {
	DB *sql.DB
}

func (r *PostgresUser) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query("Select id,first_name,last_name,created_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *PostgresUser) CreateUser(user models.User) (models.User, error) {
	query := `insert into users(first_name,last_name,created_at) values ($1,$2,$3) returning id`
	err := r.DB.QueryRow(query, &user.FirstName, &user.LastName, time.Now()).Scan(&user.ID)
	if err != nil {
		return user, err
	}
	user.CreatedAt = time.Now()
	return user, nil
}
