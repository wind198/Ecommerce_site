package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//todo
	//get user data from request body
	//create a userdata struct and decode request body into it
	//connect to database and check if a user with such email already exist
	//if not exist add it to db and anwser to response
	db := newDbHanlder.db
	var userToRegister userToCheck
	err := json.NewDecoder(r.Body).Decode(&userToRegister)
	log.Println(userToRegister)
	if err != nil {
		log.Println("Error during decoding request body", err)
		return
	}
	q := `select userEmail from user where userEmail = ?`
	row := db.QueryRow(q, userToRegister.UserEmail)
	var userAlreadyExist userToCheck
	err = row.Scan(&userAlreadyExist.UserEmail)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Conten-Type", "application/text")
		w.Write([]byte("This email is already exist"))
		return
	} else if err != sql.ErrNoRows {
		log.Println("Error during checking Email against existing ones")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sorry, we cannot check your email right now, please come back later"))
		return
	}

	q = `insert into user(userEmail,userPassword)
	values
	(?,?);
	`
	res, err := db.Exec(q, &userToRegister.UserEmail, &userToRegister.UserPassword)
	if err != nil {
		log.Println("Error during inserting to db, ", err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sorry, we cannot check your email right now, please come back later"))
	}
	log.Println(res.RowsAffected())
	w.WriteHeader(http.StatusOK)
}
