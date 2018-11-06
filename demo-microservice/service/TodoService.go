package service

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Joby81/golang/demo-microservice/data"  // NEW
	
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	type Todos []data.Todo

    todos := Todos{
        data.Todo{Name: "Write presentation"},
        data.Todo{Name: "Host meetup"},
    }

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}
