package controller

import (
    "github.com/Joby81/golang/demo-microservice/service" // NEW
    "github.com/gorilla/mux"
    "github.com/Joby81/golang/golang-utils/goutils"

)


func GetTodoRoutes() *mux.Router {

    var index = goutils.Route{
        "Index",
        "GET",
        "/",
        service.Index,
    };

    var todoIndex= goutils.Route{
        "TodoIndex",
        "GET",
        "/todos",
        service.TodoIndex,
    }

    var todoShow=goutils.Route{
        "TodoShow",
        "GET",
        "/todos/{todoId}",
        service.TodoShow,
    }

    goutils.CreateNewRoute(index);
    goutils.CreateNewRoute(todoIndex);
    goutils.CreateNewRoute(todoShow);
    return goutils.GetRoutes()
}

