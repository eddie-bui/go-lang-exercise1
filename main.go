package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	// "log"

    // a Go HTTP package for handling HTTP requesting when creating GO APIs
	"net/http"

    // for URL matcher and routing
	// "github.com/gorilla/mux"

    // a Go PostgreSQL driver for handling the database/SQL package.
	_ "github.com/lib/pq"
    "go-lang-exercise1/package/test"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "12345678x@X"
    DB_NAME     = "Exercise"
)

// Define table struct
type User struct {
    ID   int `json:"id"`
    userName string `json:"username"`
    email string `json:"email"`
    phone int `json:"phone"`
    DOB string `json:"dateOfBirth"`
}
// Define Response struct
type JsonResponse struct {
    Type    string `json:"type"`
    Data    []User `json:"data"`
    Message string `json:"message"`
}

func checkErr(err error) {
    if err != nil {
        fmt.Println("Error: ", err)
        panic(err)
    }
}

func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}

// DB set up
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return db
}

// Get all users

// response and request handlers
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    db := setupDB()

    printMessage("Getting users...")

    // Get all users from Users table 
    rows, err := db.Query("SELECT * FROM users")

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var users []User

    // Foreach movie
    for rows.Next() {
        var id int
        var username string
        var phone int
        var email string
        var dateOfBirth string

        err = rows.Scan(&id, &username, &email, &phone, &dateOfBirth)

        // check errors
        checkErr(err)

        users = append(users, User{ ID: id, userName: username, phone: phone, email: email, DOB: dateOfBirth})
    }

    var response = JsonResponse{Type: "success", Data: users}

    json.NewEncoder(w).Encode(response)
}

func main() {
    test.Hello()
	// Init the mux router
    // router := mux.NewRouter()

    // // Route handles & endpoints

    // // Get all users
    // router.HandleFunc("/api/users", GetAllUsers).Methods("GET")
    
    // // Get users by id
    // // router.HandleFunc("/users/{usersid}", GetUserById).Methods("GET")

    // // Create a users
    // // router.HandleFunc("/users/", CreateUser).Methods("POST")

    // // Update a specific users by the usersID
    // // router.HandleFunc("/users/{usersid}", UpdateUser).Methods("PUT")
    
    // // Delete a specific users by the usersID
    // // router.HandleFunc("/users/{usersid}", DeleteUser).Methods("DELETE")

    
    // // serve the app)
    // fmt.Println("Server at 8080")
    // log.Fatal(http.ListenAndServe(":8000", router))

    
}