package main

import (
    "net/http"
    //"database/sql"

    //"github.com/gorilla/mux"
    //_ "github.com/lib/pq"
)

type User struct {
    Id              int       `json:"id"`
    FirstName       string    `json:"first_name"`
    LastName        string    `json:"last_name"`
    Email           string    `json:"email"`
    Active          bool      `json:"active"`
}

type Users []User

func GetUsers(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}

func GetUser(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}

func PostUser(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}
