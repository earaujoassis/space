package main

import (
    "net/http"

    . "github.com/earaujoassis/space/energy"
)

const (
    GET    = "GET"
    POST   = "POST"
    PUT    = "PUT"
    DELETE = "DELETE"
)

func Index(r *http.Request) (int, interface{}) {
    return 200, map[string]string{"message": "Hello World"}
}

var routes = Routes{
    Route{ GET, "/", Index, },
    Route{ GET, "/applications", GetApplications, },
    Route{ POST, "/applications", PostApplication, },
    Route{ GET, "/applications/{id}", GetApplication, },
    Route{ GET, "/users", GetUsers, },
    Route{ POST, "/users", PostUser, },
    Route{ GET, "/users/{id}", GetUser, },
}
