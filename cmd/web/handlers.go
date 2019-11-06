package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ECAllen/lets-go/pkg/forms"
	"github.com/ECAllen/lets-go/pkg/models"
)

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sign up the user...")
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display login user form...")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login user...")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout user...")
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	m, err := app.memories.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Memories: m})
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

	app.render(w, r, "show.page.tmpl", &templateData{Memory: mid})
}

func (app *application) createMemoryForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createMemory(w http.ResponseWriter, r *http.Request) {

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
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	id, err := app.memories.Insert(form.Get("title"), form.Get("content"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Log successfully created")
	http.Redirect(w, r, fmt.Sprintf("/memory/%d", id), http.StatusSeeOther)
}
