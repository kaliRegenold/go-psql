package main

import(
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
)

type user struct {
    UserId      int
    Lastname    string
    Firstname   string
    State       string
    City        string
}

func error_out(errCode int, w http.ResponseWriter, message string) {
    w.WriteHeader(errCode)
    w.Write([]byte(message))
}


func get_user(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  u := user{}
  stmt := `SELECT user_id, lastname, firstname, state, city FROM "users" WHERE user_id=$1;`

  err := db.QueryRow(stmt, params["id"]).Scan(
    &u.UserId,
    &u.Lastname,
    &u.Firstname,
    &u.State,
    &u.City,)

  if err != nil {
    error_out(404, w, fmt.Sprintf("%s\n", err))
    return
  }

  json_user, err := json.Marshal(u)

  if err != nil {
    error_out(500, w, fmt.Sprintf("%s\n", err))
    return
  }

  w.WriteHeader(200)
  w.Write(json_user)
}
