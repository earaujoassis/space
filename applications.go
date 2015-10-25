package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

type Application struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func (a Application) Get(id string) {

}

type Applications []Application

func (a Applications) Find() {

}

func GetApplications(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}

func GetApplication(r *http.Request) (int, interface{}) {
    var application Application
    vars := mux.Vars(r)
    id := vars["id"]

    application.Get(id)
    return 200, application
}

func PostApplication(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}
