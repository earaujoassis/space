package main

import (
    "net/http"

    "./energy"
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

var routes = energy.Routes{
    energy.Route{ GET, "/", Index, },
    energy.Route{ GET, "/applications", GetApplications, },
    energy.Route{ POST, "/applications", PostApplication, },
    energy.Route{ GET, "/applications/{id}", GetApplication, },
    energy.Route{ GET, "/users", GetUsers, },
    energy.Route{ POST, "/users", PostUser, },
    energy.Route{ GET, "/users/{id}", GetUser, },
}
