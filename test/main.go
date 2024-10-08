package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func upHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "OK")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/up", upHandler)

    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Server failed: %s\n", err)
    }
}
