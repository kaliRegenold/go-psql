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
    r.HandleFunc("/ping", ping_handler).Methods("GET")
    r.HandleFunc("/user/{id}", get_user).Methods("GET")
    http.ListenAndServe(":8080", r)
}


func ping_handler(w http.ResponseWriter, r *http.Request) {
    err := db.Ping()
    if err != nil {
        fmt.Fprintf(w, "%s\n", err)
    } else {
        fmt.Fprintf(w, "Success pinging database\n")
    }
}
