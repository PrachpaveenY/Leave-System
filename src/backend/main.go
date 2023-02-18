package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Welcome to Server")

	InitDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// login
	http.HandleFunc("/logins", loginHandler)
	http.HandleFunc("/loginauths", loginAuthHandler)
	http.HandleFunc("/registers", registersHandler)
	http.HandleFunc("/registerauths", registerauthsHandler)

	log.Fatal(e.Start(":2566"))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) error {
	start := time.Now()
	fmt.Println("loginHandler running!!!")
	time.Sleep(1 * time.Second)
	fmt.Println("Run Time : ", time.Since(start), "sec")
	tpl.ExecuteTemplate(w, "/src/main.js", nil)
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("loginAuthHandler running!!!")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username:", username, "password:", password)
	//
	var login string
	stmt := 
	row := db.QueryRow(stmt, username)
	err := row.Scan(&login)
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println("login from data:", login)
		tpl.ExecuteTemplate(w, "/src/main.js", "check username and password")
		return
	}
	//
	err := bcrypt.CompareHashAndPassword([]byte(login), []byte(password))
	if err == nil {
		fmt.Fprint(w, "successfully Login")
		return
	}
	fmt.Println("incorrect password")
	tpl.ExecuteTemplate(w, "/src/main.js", check username and password)
}

func registersHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("registersHandler running!!!")
	time.Sleep(1 * time.Second)
	if err != nil {

		return
	}
}

func registerauthsHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("registerauthsHandler running!!!")
	time.Sleep(1 * time.Second)
	if err != nil {

		return
	}
}

// close connection
var db *sql.DB
var tpl *template.Template

// Get database url from environment variable, Create Table
func InitDB() {
	tpl, _ = template.ParseGlob("")
	var err error
	// connect to MongoDB
	db, err := sql.Open("mongodb", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	createLogin := `
	CREATE TABLE IF NOT EXISTS logins (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[] );
	`
	_, err = db.Exec(createLogin)
	if err != nil {
		log.Fatal("can't create table login", err)
	}

	createUser := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[] );
	`
	_, err = db.Exec(createUser)
	if err != nil {
		log.Fatal("can't create table user", err)
	}
}

// ===========

// connect to MongoDB
// db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
// if err != nil {
// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 	os.Exit(1)
// }
// defer db.Close()

// 	var greeting string
// 	err = db.QueryRow("select 'Hello'").Scan(&greeting)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(greeting)
