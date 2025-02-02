package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mmkader85/build-jwt-authenticated-restful-apis-with-golang/controllers"
	"github.com/mmkader85/build-jwt-authenticated-restful-apis-with-golang/driver"
	"github.com/mmkader85/build-jwt-authenticated-restful-apis-with-golang/utils"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db = driver.ConnectDB()
	controller := controllers.Controller{}

	r := mux.NewRouter()
	r.HandleFunc("/signup", controller.SignUpHandler(db)).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler(db)).Methods("POST")
	r.HandleFunc("/get_all_users", utils.TokenVerifyMiddleware(controller.GetAllUsers(db))).Methods("GET")

	fmt.Println("Listening and serving at port 8000...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))
}
