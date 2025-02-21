# Simple Go (Golang) REST API (JSON)

The task requires implementing a GET API endpoint in Go using the `gorilla/mux` router. 
The endpoint `/users` should return user data from a mocked database and support filtering by the `name` query parameter.

## Installation & Usage

1. Setup a Go project:
```
go mod init myapi
go get github.com/gorilla/mux

```

2. Create `main.go`:
```
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
```

3. Result:

- `GET /users` → Returns all users.
- `GET /users?name=John` → Returns only users with `name="John"`.
- If no user matches the filter, returns an empty array `[]`.
