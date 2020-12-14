package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mmkader85/build-jwt-authenticated-restful-apis-with-golang/models"
	userRepository "github.com/mmkader85/build-jwt-authenticated-restful-apis-with-golang/repository/user"
	"github.com/mmkader85/build-jwt-authenticated-restful-apis-with-golang/utils"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct{}

func (c Controller) SignUpHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sign-up reached...")
		var user models.User
		var signUpErr models.Error
		var userRepo userRepository.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			signUpErr.Message = "Bad signup request!"
			utils.RespondWithError(w, http.StatusBadRequest, signUpErr)
			return
		}

		if strings.TrimSpace(user.Email) == "" {
			signUpErr.Message = "Email is missing!"
			utils.RespondWithError(w, http.StatusBadRequest, signUpErr)
			return
		}

		if strings.TrimSpace(user.Password) == "" {
			signUpErr.Message = "Password is missing!"
			utils.RespondWithError(w, http.StatusBadRequest, signUpErr)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			signUpErr.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, signUpErr)
			return
		}

		user.Password = string(hash)
		err = userRepo.CreateUser(db, &user)
		if err != nil {
			signUpErr.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, signUpErr)
			return
		}

		utils.ResponseJson(w, http.StatusOK, struct {
			ID    int
			Email string
		}{
			ID:    user.ID,
			Email: user.Email,
		})
	}
}

func (c Controller) LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login reached...")

		var jwt models.JWT
		var user models.User
		var loginErr models.Error
		var userRepo userRepository.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			loginErr.Message = "Bad login request!"
			utils.RespondWithError(w, http.StatusBadRequest, loginErr)
			return
		}

		inputPwd := user.Password
		hashedDbPwd, userErr := userRepo.GetPasswordByEmail(db, user.Email)
		if userErr == sql.ErrNoRows {
			loginErr.Message = "Email doesn't exist"
			utils.RespondWithError(w, http.StatusBadRequest, loginErr)
			return
		} else if userErr != nil {
			loginErr.Message = "Database connection error"
			utils.RespondWithError(w, http.StatusInternalServerError, loginErr)
			return
		}

		pwdErr := bcrypt.CompareHashAndPassword([]byte(hashedDbPwd), []byte(inputPwd))
		if pwdErr != nil {
			loginErr.Message = "Incorrect password"
			utils.RespondWithError(w, http.StatusBadRequest, loginErr)
			return
		}

		tokenString, tErr := utils.GenerateToken(user)
		if tErr != nil {
			loginErr.Message = tErr.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, loginErr)
			return
		}
		jwt.Token = tokenString

		utils.ResponseJson(w, http.StatusOK, jwt)
	}
}
