package main

import (
    "github.com/gorilla/mux"
    "database/sql"
    "net/http"
)

var p_db *sql.DB

func main() {
    // Postgres
    p_db = p_postgres_init()
    defer p_db.Close()

    // Redis
    r_redis_init()

    r := mux.NewRouter()
    r.Use(p_common_middleware)
    r.HandleFunc("/ping", p_ping_handler).Methods("GET")
    r.HandleFunc("/user/{id}", get_user).Methods("GET")
    http.ListenAndServe(":8080", r)
}
