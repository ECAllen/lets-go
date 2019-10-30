package main

import (
	"net/http"
	"fmt"
	"strconv"
	"errors"
	"github.com/ECAllen/lets-go/pkg/models"
	"github.com/ECAllen/lets-go/pkg/forms"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	m, err := app.memories.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Memories: m,})
}

func (app *application) showMemory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}


	mid, err := app.memories.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Memory: mid,})
}

func (app *application) createMemoryForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createMemory(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w,r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	id, err := app.memories.Insert(form.Get("title"),form.Get("content"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/memory/%d", id), http.StatusSeeOther)
}


