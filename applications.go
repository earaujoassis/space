package main

import (
    "net/http"
    //"database/sql"

    //"github.com/gorilla/mux"
    //_ "github.com/lib/pq"
)

type Application struct {
    Id              int       `json:"id"`
    Name            string    `json:"name"`
}

type Applications []Application

func GetApplications(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}

func GetApplication(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}

func PostApplication(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}
