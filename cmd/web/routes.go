package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showMemory)
	mux.HandleFunc("snippet/create", app.createMemory)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}