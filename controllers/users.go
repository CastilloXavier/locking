package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"lenslocked.com/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the template are not parsed
// correctly, and should only be used during initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

//New is used to render the form where a user can create a new user account
//
//GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//Create is used to process the signup form when a users
//Submit it. This is used to create a new user account
//
//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	var signupForm SignupForm
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&signupForm, r.PostForm); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, signupForm)
}
