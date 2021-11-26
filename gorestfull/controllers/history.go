package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
)

type histories struct {
	rp interfaces.HistRepository
}

func History(rps interfaces.HistRepository) *histories {
	return &histories{rps}
}

func (hist *histories) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	data, err := hist.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Response(w, &data, 200, false)
}

func (hist *histories) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	var body models.History

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	data := models.CreateHistory()
	data.Id = body.Id
	data.No_invoice = body.No_invoice
	data.Cashier = body.Cashier
	data.Orders = body.Orders
	data.Amount = body.Amount

	err = hist.rp.Save(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The history has been successfully saved"))
	helpers.Response(w, &data, 201, false)
}

func (hist *histories) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	var body models.History
	vars := mux.Vars(r)
	historyID := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	data := models.CreateHistory()
	data.No_invoice = body.No_invoice
	data.Cashier = body.Cashier
	data.Orders = body.Orders
	data.Amount = body.Amount

	err = hist.rp.Update(data, historyID)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The history has been successfully updated"))
	helpers.Response(w, &data, 200, false)
}

func (hist *histories) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	vars := mux.Vars(r)
	historyID := vars["id"]

	err := hist.rp.Remove(historyID)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The history has been successfully deleted"))
	helpers.Response(w, "The history has been successfully deleted", 201, false)
}
