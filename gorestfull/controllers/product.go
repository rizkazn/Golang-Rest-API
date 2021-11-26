package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
)

type products struct {
	rp interfaces.ProdRepository
}

func Products(rps interfaces.ProdRepository) *products {
	return &products{rps}
}

func (prod *products) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	data, err := prod.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Response(w, &data, 200, false)
}

func (prod *products) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	var body models.Product

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	data := models.CreateProduct()
	data.Id = body.Id
	data.Product_name = body.Product_name
	data.Price = body.Price
	data.Product_image = body.Product_image
	data.Category_id = body.Category_id

	err = prod.rp.Save(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The product has been successfully saved"))
	helpers.Response(w, &data, 201, false)
}

func (prod *products) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	var body models.Product
	vars := mux.Vars(r)
	productID := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	data := models.CreateProduct()
	data.Product_name = body.Product_name
	data.Price = body.Price
	data.Category_id = body.Category_id

	err = prod.rp.Update(data, productID)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The product has been successfully updated"))
	helpers.Response(w, &data, 200, false)
}

func (prod *products) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	vars := mux.Vars(r)
	productId := vars["id"]

	err := prod.rp.Remove(productId)

	if err != nil {
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("The product has been successfully deleted"))
	helpers.Response(w, "The product has been successfully deleted", 200, false)
}

func (prod *products) SearchByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	// vars := r.URL.Query()
	// catchName := vars["product_name"]
	// productName := strings.Join(catchName, "")
	// data, err := prod.rp.SearchProducBytName(productName)

	vars := mux.Vars(r)

	productName := vars["product_name"]

	data, err := prod.rp.SearchProductBytName(productName)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Response(w, &data, 200, false)
}

func (prod *products) SortByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	// vars := mux.Vars(r)
	// productName := vars["product_name"]

	data, err := prod.rp.OrderProdByName()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Response(w, &data, 200, false)
}

func (prod *products) SortByCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	// vars := mux.Vars(r)
	// categoryId := vars["category_id"]

	data, err := prod.rp.OrderProdByCategory()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Response(w, &data, 200, false)
}

func (prod *products) SortByNewest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	// vars := mux.Vars(r)
	// updatedAt := vars["updated_at"]

	data, err := prod.rp.OrderProdByNewest()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Response(w, &data, 200, false)
}

func (prod *products) SortByPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	// vars := r.URL.Query()
	// catchName := vars["price"]
	// price := strings.Join(catchName, " ")
	// data, err := prod.rp.OrderProdByPrice(price)

	// vars := mux.Vars(r)
	// productPrice := vars["price"]
	// data, err := prod.rp.OrderProdByPrice(productPrice)

	data, err := prod.rp.OrderProdByPrice()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Response(w, &data, 200, false)
}
