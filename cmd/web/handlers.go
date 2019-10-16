package main

import (
	"net/http"
	"fmt"
	"strconv"
	"errors"
	"github.com/ECAllen/lets-go/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	m, err := app.memories.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, memory := range m {
		fmt.Fprintf(w, "%v\n", memory)
	}

//	files := []string{
//		"./ui/html/home.page.tmpl",
//		"./ui/html/base.layout.tmpl",
//		"./ui/html/footer.partial.tmpl",
//	}
//
//	ts, err := template.ParseFiles(files...)
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//
//	err = ts.Execute(w, nil)
//	if err != nil {
//		app.serverError(w, err)
//	}

}

func (app *application) showMemory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}


	m, err := app.memories.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%v", m)
}


func (app *application) createMemory(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"

	id, err := app.memories.Insert(title,content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("memory?id=%id", id), http.StatusSeeOther)
}
