package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"udemy/build-jwt-authenticated-restful-apis-with-golang/models"
	"udemy/build-jwt-authenticated-restful-apis-with-golang/utils"
)

func (c Controller) GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Protected reached...")
		var err models.Error
		var users []string

		stmt := "SELECT id, email FROM users ORDER BY id DESC LIMIT 10;"
		userRows, error := db.Query(stmt)
		if error == sql.ErrNoRows {
			err.Message = "No users found!"
			utils.RespondWithError(w, http.StatusOK, err)
			return
		} else if error != nil {
			err.Message = error.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		defer userRows.Close()

		for userRows.Next() {
			var id int
			var email string
			resultErr := userRows.Scan(&id, &email)
			if resultErr != nil {
				err.Message = resultErr.Error()
				utils.RespondWithError(w, http.StatusInternalServerError, err)
				return
			}
			users = append(users, email)
		}
		utils.ResponseJson(w, http.StatusOK, users)
		//utils.ResponseJson(w, http.StatusOK, struct {
		//	Status string
		//}{
		//	"Ok",
		//})
	}
}
