package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/mmkader85/build-jwt-authenticated-restful-apis-with-golang/models"
	"github.com/mmkader85/build-jwt-authenticated-restful-apis-with-golang/utils"
)

func (c Controller) GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Protected reached...")
		var err models.Error
		var users []string

		stmt := "SELECT id, email FROM users ORDER BY id DESC LIMIT 10;"
		userRows, getUserErr := db.Query(stmt)
		if getUserErr == sql.ErrNoRows {
			err.Message = "No users found!"
			utils.RespondWithError(w, http.StatusOK, err)
			return
		} else if getUserErr != nil {
			err.Message = getUserErr.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		defer func() {
			_ = userRows.Close()
		}()

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
	}
}
