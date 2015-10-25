package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

type User struct {
    Id        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Active    bool   `json:"active"`
}

func (u User) Get(id string) {

}

type Users []User

func (u Users) Find() {

}

func GetUsers(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}

func GetUser(r *http.Request) (int, interface{}) {
    var user User
    vars := mux.Vars(r)
    id := vars["id"]

    user.Get(id)
    return 200, user
}

func PostUser(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}
