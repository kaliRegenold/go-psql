package main

import (
    "fmt"
    "os"
    _ "github.com/lib/pq"
    "database/sql"
    "net/http"
)

func p_postgres_init() *sql.DB {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
        os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
        os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"))
    d, err := sql.Open("postgres", psqlInfo)

    if err != nil {
        panic("Error opening database\n")
    }

    return d
}

func p_common_middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func p_ping_handler(w http.ResponseWriter, r *http.Request) {
    err := p_db.Ping()
    if err != nil {
        error_out(500, w, fmt.Sprintf("%s\n", err))
    } else {
        w.WriteHeader(200)
        w.Write([]byte("Success pinging database\n"))
    }
}
