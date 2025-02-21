package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/codility/rest_api_go/database"
)

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := database.GetUsers()
    queryName := r.URL.Query().Get("name")

    var filteredUsers []database.User
    for _, user := range users {
        if queryName == "" || user.Name == queryName {
            filteredUsers = append(filteredUsers, user)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(filteredUsers)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/users", getUsersHandler).Methods("GET")

    http.ListenAndServe(":8080", r)
}
