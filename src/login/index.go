package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Welcome to Server")

	InitDB()
	e := echo.New()
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "mongodb" || password == "2566" {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// login
	// e.POST("/logins", CreateLoginsAllHandler)
	// e.GET("/logins", GetLoginsHandler)
	// e.GET("/logins/:id", GetLoginsIDHandler)
	// e.PUT("/logins/:id", UpdateAllLoginsHandler)
	// e.PATCH("/logins/:id", UpdateLoginsHandler)
	// e.DELETE("/logins/:id", DeleteLoginsHandler)

	// user
	// e.POST("/users", CreateUsersAllHandler)
	// e.GET("/users", GetUsersHandler)
	// e.GET("/users/:id", GetUsersIDHandler)
	// e.PUT("/users/:id", UpdateAllUsersHandler)
	// e.PATCH("/users/:id", UpdateUsersHandler)
	// e.DELETE("/users/:id", DeleteUsersHandler)

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
