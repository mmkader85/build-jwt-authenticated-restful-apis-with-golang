package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"udemy/build-jwt-authenticated-restful-apis-with-golang/models"
	userRepository "udemy/build-jwt-authenticated-restful-apis-with-golang/repository/user"
	"udemy/build-jwt-authenticated-restful-apis-with-golang/utils"
)

type Controller struct{}

func (c Controller) SignUpHandler(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sign-up reached...")
		var user models.User
		var error models.Error
		var userRepo userRepository.User

		json.NewDecoder(r.Body).Decode(&user)

		if strings.TrimSpace(user.Email) == "" {
			error.Message = "Email is missing!"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if strings.TrimSpace(user.Password) == "" {
			error.Message = "Password is missing!"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			error.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		user.Password = string(hash)
		err = userRepo.CreateUser(db, &user)
		if err != nil {
			error.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.ResponseJson(w, http.StatusOK, struct {
			ID    int
			Email string
		}{
			ID:    user.ID,
			Email: user.Email,
		})

		//stmt := "SELECT * FROM users WHERE email = $1;"
		//queryErr := db.QueryRow(stmt, user.Email).Scan(&user.ID, &user.Email, &user.Password)
		//if queryErr != nil {
		//	error.Message = queryErr.Error()
		//	utils.RespondWithError(w, http.StatusInternalServerError, error)
		//	return
		//}
		//
		//utils.ResponseJson(w, http.StatusOK, user)
	}
}

func (c Controller) LoginHandler(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login reached...")

		var jwt models.JWT
		var user models.User
		var err models.Error
		var userRepo userRepository.User

		json.NewDecoder(r.Body).Decode(&user)
		inputPwd := user.Password

		hashedDbPwd, userErr := userRepo.GetPasswordByEmail(db, user.Email)
		if userErr == sql.ErrNoRows {
			err.Message = "Email doesn't exist"
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		} else if userErr != nil {
			err.Message = "Database connection error"
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		pwdErr := bcrypt.CompareHashAndPassword([]byte(hashedDbPwd), []byte(inputPwd))
		if pwdErr != nil {
			err.Message = "Incorrect password"
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		tokenString, tErr := utils.GenerateToken(user)
		if tErr != nil {
			err.Message = tErr.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		jwt.Token = tokenString

		utils.ResponseJson(w, http.StatusOK, jwt)
	}
}