package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

    // a Go HTTP package for handling HTTP requesting when creating GO APIs
	"net/http"

    // for URL matcher and routing
	"github.com/gorilla/mux"

    // a Go PostgreSQL driver for handling the database/SQL package.
	_ "github.com/lib/pq"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "12345678x@X"
    DB_NAME     = "Exercise"
)

// Define table struct
type User struct {
    ID   string 
    userName string 
    email string 
    phone int 
    DOB string 
}
// Define Response struct
type JsonResponse struct {
    Type    string 
    Data    []User 
    Message string 
}
// DB set up
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return db
}

func main() {
	test.Hello()
	// do(21)
	// do("hello")
	// do(true)
}