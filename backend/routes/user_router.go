package routes

import (
	"api/controllers"
	"database/sql"

	"github.com/gorilla/mux"
)

func UserRoutes(db *sql.DB) {
	router := mux.NewRouter()
  router.HandleFunc("/api/v1/go/users", controllers.GetUsers(db)).Methods("GET")
  router.HandleFunc("/api/v1/go/users", controllers.CreateUser(db)).Methods("POST")
  router.HandleFunc("/api/v1/go/users/{id}", controllers.GetUser(db)).Methods("GET")
  router.HandleFunc("/api/v1/go/users", controllers.UpdateUser(db)).Methods("PUT")
  router.HandleFunc("/api/v1/go/users", controllers.DeleteUser(db)).Methods("DELETE")
}