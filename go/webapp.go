package main

import (
    _ "github.com/lib/pq"
    "database/sql"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "os"
)

var db *sql.DB

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
        os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
        os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"))
    d, err := sql.Open("postgres", psqlInfo)
    defer d.Close()

    if err != nil {
        panic("Error opening database\n")
    }

    db = d

    r := mux.NewRouter()
    r.Use(commonMiddleware)
    r.HandleFunc("/ping", ping_handler).Methods("GET")
    r.HandleFunc("/user/{id}", get_user).Methods("GET")
    http.ListenAndServe(":8080", r)
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func ping_handler(w http.ResponseWriter, r *http.Request) {
    err := db.Ping()
    if err != nil {
        error(500, w, fmt.Sprintf("%s\n", err))
    } else {
        w.WriteHeader(200)
        w.Write([]byte("Success pinging database\n"))
    }
}
