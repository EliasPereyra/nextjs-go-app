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

func CreateUser(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		json.NewDecoder(r.Body).Decode(&u)

		err := db.QueryRow("INSERT INTO users (fullname, email, profile_img) VALUES ($1, $2, $3) RETURNING id", u.Fullname, u.Email, u.Profile_img).Scan(&u.Id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	} 
}

func UpdateUser(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		json.NewDecoder(r.Body).Decode(&u)

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE users SET fullname = $1, email = $2, profile_img = $3 WHERE id = $4", u.Fullname, u.Email, u.Profile_img, id)
		if err != nil {
			log.Fatal(err)
		}

		var updatedUser model.User
		err = db.QueryRow("SELECT id, fullname, email, profile_img FROM users WHEERE id = $1", id).Scan(&updatedUser.Id, &updatedUser.Fullname, &updatedUser.Email, &updatedUser.Profile_img)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(updatedUser)
	}
}

func DeleteUser(db *sql.DB){}