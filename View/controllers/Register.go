package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	// "io"
)

type User struct {
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Dob   string `json:"Dob"`
}

var db *sql.DB

func PostForm(w http.ResponseWriter, r *http.Request) {
	// Only POST routes allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
		return
	}

	var user User
	//Load ENV
	godotenv.Load()

	// Connect config for mysql DB
	conn := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB"),
	}

	var error error
	// Connect to DB
	db, error = sql.Open("mysql", conn.FormatDSN())
	if error != nil {
		http.Error(w, "Could not connect to the database", http.StatusBadRequest)
	}

	//Read Body data
	data, dataError := io.ReadAll(r.Body)
	if dataError != nil {
		http.Error(w, "Could not read body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	//Creates a buffer for body data
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// decodes & stores data into user
	jsonErr := json.NewDecoder(r.Body).Decode(&user)

	if jsonErr != nil {
		http.Error(w, "Could not create user", http.StatusBadRequest)
	}

	//Insert data into DB
	dbExec, dbError := db.Exec("INSERT INTO users (Fname, Lname, dob) VALUES (?,?,?)", user.Fname, user.Lname, user.Dob)
	if dbError != nil {
		http.Error(w, "Could not submit user", http.StatusBadRequest)
		fmt.Printf("dbError: %v", dbError)
	}
	if dbExec != nil {
		http.Error(w, "Success", http.StatusAccepted)
	}
}
