package goutils

import (
    "github.com/gorilla/mux"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}


func GetRoutes() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
    return router
}

type Routes []Route
var routes = Routes{}

func CreateNewRoute(route Route)  {
    routes = append(routes, route);
}