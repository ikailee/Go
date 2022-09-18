package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
		},
	}

	temp.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	temp = template.Must(template.ParseFiles("templates/index.html"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":5000", mux))
}
