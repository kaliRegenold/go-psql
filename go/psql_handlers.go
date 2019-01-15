package main

import(
    // "encoding/json"
    // "github.com/gorilla/mux"
    "net/http"
    // "fmt"
)

type user struct {
    user_id     int
    lastname    string
    firstname   string
    state       string
    city        string
}

type users struct {
}

func get_all_user(w http.ResponseWriter, r *http.Request) {

}

func get_user(w http.ResponseWriter, r *http.Request) {
    // params := mux.Vars(r)
    // u := user{}
    // // stmt := `SELECT
    // //             user_id,
    // //             lastname,
    // //             firstname,
    // //             state,
    // //             city
    // //         FROM users
    // //         WHERE user_id=$1;`
    // stmt := `SELECT * FROM users WHERE user_id=$1;`
    // 
    // err := db.QueryRow(stmt, params["id"]).Scan(
    //         &u.user_id,
    //         &u.lastname,
    //         &u.firstname,
    //         &u.state,
    //         &u.city,)
    // 
    // if err != nil {
    //     fmt.Fprintf(w, "%s\n", err)
    //     return
    // }
    // 
    // json_user, err := json.Marshal(u)
    // if err != nil {
    //     fmt.Fprintf(w, "%s\n", err)
    //     return
    // }
    // 
    // fmt.Fprintf(w, string(json_user))
}

func create_user(w http.ResponseWriter, r *http.Request) {

}

func delete_user(w http.ResponseWriter, r *http.Request) {

}
