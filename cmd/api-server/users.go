package main

import (
	"net/http"

	"github.com/joddyhartono/goapi/internal/models"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload
	err := readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &models.User{
		Username: payload.Username,
		Email: payload.Email,
		Password: payload.Password,
	}

	ctx := r.Context()

	err = app.store.Users.Create(ctx, user)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusCreated, user)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}