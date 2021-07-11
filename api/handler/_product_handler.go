package handler

import (
	"net/http"
	"salon-booking-guru/store/model"
)

func productRoutes() {
	router.HandleFunc("/product", getAllProduct).Methods("GET")
	router.HandleFunc("/product/{id}", getProduct).Methods("GET")
	router.HandleFunc("/product", createProduct).Methods("POST")
	router.HandleFunc("/product", updateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
}

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := s.Product().GetAll()
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	respond(w, products, http.StatusOK)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	product, err := s.Product().Get(id)
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve product", http.StatusInternalServerError)
		return
	}

	respond(w, product, http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := readBytes(w, r, &product)
	if err != nil {
		return
	}

	err = s.Product().Create(&product)
	if err != nil {
		respondMsg(w, "Error: Failed to create product", http.StatusInternalServerError)
		return
	}

	respond(w, product, http.StatusOK)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := readBytes(w, r, &product)
	if err != nil {
		return
	}

	err = s.Product().Update(&product)
	if err != nil {
		respondMsg(w, "Error: Failed to update product", http.StatusInternalServerError)
		return
	}

	respond(w, product, http.StatusOK)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	err = s.Product().Delete(id)
	if err != nil {
		respondMsg(w, "Error: Failed to delete product", http.StatusInternalServerError)
		return
	}

	respondEmpty(w, http.StatusOK)
}
