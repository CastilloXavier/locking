package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/models"
	"lenslocked.com/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the template are not parsed
// correctly, and should only be used during initial setup.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

type Users struct {
	NewView *views.View
	us      *models.UserService
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
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//Create is used to process the signup form when a users
//Submit it. This is used to create a new user account
//
//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var signupForm SignupForm
	if err := parseForm(r, &signupForm); err != nil {
		panic(err)
	}
	user := models.User{
		Name:     signupForm.Name,
		Email:    signupForm.Email,
		Password: signupForm.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, user)
}
