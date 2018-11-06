package main

import (
    "log"
    "net/http"
    "github.com/Joby81/golang/demo-microservice/controller"  // NEW	
)


func main() {
    router := controller.GetTodoRoutes()
    log.Fatal(http.ListenAndServe(":8080", router))
}

