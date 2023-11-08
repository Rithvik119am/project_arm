package main

import (
	"net/http"
	"text/template"
)

func main() {
	/*postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	//this are changes
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
