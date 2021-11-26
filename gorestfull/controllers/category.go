package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
)

type categories struct {
	rp interfaces.CatRepository
}

func Category(rps interfaces.CatRepository) *categories {
	return &categories{rps}
}

func (cat *categories) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")
	data, err := cat.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Response(w, &data, 200, false)
}

func (cat *categories) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")
	var body models.Category

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	data := models.CreateCategory()
	data.Id = body.Id
	data.Category_name = body.Category_name
	data.Category_image = body.Category_image

	err = cat.rp.Save(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The category has been successfully saved"))
	helpers.Response(w, &data, 201, false)
}

func (cat *categories) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	var body models.Category
	vars := mux.Vars(r)
	categoryID := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	data := models.CreateCategory()
	data.Category_name = body.Category_name

	err = cat.rp.Update(data, categoryID)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The category has been successfully updated"))
	helpers.Response(w, &data, 200, false)
}

func (cat *categories) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	vars := mux.Vars(r)
	categoryID := vars["id"]

	err := cat.rp.Remove(categoryID)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The category has been successfully deleted"))
	helpers.Response(w, "The category has been successfully deleted", 200, false)
}
