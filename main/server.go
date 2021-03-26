package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	asciiart ".."
)

var templates = template.Must(template.ParseFiles("../static/index.html"))

func createASCII(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	text := r.FormValue("text")
	allStr := strings.Split(text, "\n")
	font := r.FormValue("font")
	res := asciiart.ConvertToAscii(allStr[0], font)
	t := struct {
		Str string
	}{
		Str: res,
	}
	err := templates.ExecuteTemplate(w, "index.html", t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main() {
	fmt.Println("listening...")
	http.Handle("/", http.FileServer(http.Dir("../static")))
	http.HandleFunc("/action", createASCII)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
