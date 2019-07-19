package main

import (
    "io"
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Name string `json:"name"`
}

type Message struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        response := Message {
            Message: fmt.Sprintf("Hello %s", user.Name),
        }
        json.NewEncoder(w).Encode(response)
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "OK");
    })

    http.ListenAndServe(":8080", nil)
}
