package main

import (
	//"fmt"
	"fmt"
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
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, postgresqlDbInfo)
	})

}
