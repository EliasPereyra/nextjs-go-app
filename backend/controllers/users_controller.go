package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"api/model"

	"github.com/gorilla/mux"
)

func GetUsers(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		users := []model.User{}

		for rows.Next() {
			var u model.User
			if err := rows.Scan(&u.Id, &u.Fullname, &u.Email, &u.Profile_img); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(users)
	}
}

func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		
		var u model.User
		err := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}

func CreateUser(db *sql.DB){}

func UpdateUser(db *sql.DB){}

func DeleteUser(db *sql.DB){}