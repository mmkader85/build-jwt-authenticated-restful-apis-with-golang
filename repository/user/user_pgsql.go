package userRepository

import (
	"database/sql"
	"udemy/build-jwt-authenticated-restful-apis-with-golang/models"
)

type User struct {}

func (User) CreateUser(db *sql.DB, user *models.User) error {
	stmt := "INSERT INTO users(email, password) values ($1, $2) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (User) GetPasswordByEmail(db *sql.DB, email string) (string, error) {
	var pwd string
	stmt := "SELECT Password FROM users WHERE Email = $1;"
	err := db.QueryRow(stmt, email).Scan(&pwd)
	if err != nil {
		return "", err
	}

	return pwd, nil
}