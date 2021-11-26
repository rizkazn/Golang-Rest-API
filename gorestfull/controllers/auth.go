package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
)

type auth struct {
	rp interfaces.UserRepository
}

type userToken struct {
	Token string `json:"token"`
}

func Auth(rps interfaces.UserRepository) *auth {
	return &auth{rps}
}

func (au *auth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	var body models.User

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	password, err := au.rp.GetPassword(body.Username)

	if err != nil {
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	if !helpers.CheckPassword(password, body.Password) {
		helpers.Response(w, "Wrong Password", 401, true)
		return
	}

	role, err := au.rp.GetRole(body.Username)

	if err != nil {
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	token := helpers.GenerateToken(body.Username, role)

	tokens, err := token.Create()

	if err != nil {
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	helpers.Response(w, userToken{Token: tokens}, 200, false)
}
