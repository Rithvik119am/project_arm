package main

import (
	//"fmt"

	"html/template"
	"net/http"
)

type Team struct {
	name            string
	matches_played  int
	won             int
	lost            int
	draw            int
	goal_difference int
	points          int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "gitpod"
	password = "root"
	dbname   = "teams"
)

func main() {
	/*postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("templates/index.html"))
		tmp.Execute(w, nil)
	})
	http.HandleFunc("/getdata", func(w http.ResponseWriter, r *http.Request) {
		tmp1 := template.Must(template.ParseFiles("templates/other.html"))
		tmp1.Execute(w, nil)
	})

	http.ListenAndServe(":8000", nil)
}
