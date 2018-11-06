package controller

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Joby81/golang/demo-microservice/service"  // NEW
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}


func TodoRoutes() *mux.Router {

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

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        service.Index,
    },
    Route{
        "TodoIndex",
        "GET",
        "/todos",
        service.TodoIndex,
    },
    Route{
        "TodoShow",
        "GET",
        "/todos/{todoId}",
        service.TodoShow,
    },
}