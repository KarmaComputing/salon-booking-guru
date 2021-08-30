package handler

import (
	"net/http"
	"salon-booking-guru/store/model"
	"salon-booking-guru/validation"
)

func productCategoryRoutes() {
	// GET
	v1.HandleFunc(
		"/product-category",
		authorize(getAllProductCategory, "canReadProductCategory"),
	).Methods("GET")
	v1.HandleFunc(
		"/product-category/{id}",
		authorize(getProductCategory, "canReadProductCategory"),
	).Methods("GET")

	// POST
	v1.HandleFunc(
		"/product-category",
		authorize(createProductCategory, "canCreateProductCategory"),
	).Methods("POST")

	// PUT
	v1.HandleFunc(
		"/product-category",
		authorize(updateProductCategory, "canUpdateProductCategory"),
	).Methods("PUT")

	// DELETE
	v1.HandleFunc(
		"/product-category/{id}",
		authorize(deleteProductCategory, "canDeleteProductCategory"),
	).Methods("DELETE")
}

func getAllProductCategory(w http.ResponseWriter, r *http.Request) {
	productCategorys, err := s.ProductCategory().GetAll()
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve productCategorys", http.StatusInternalServerError)
		return
	}

	respond(w, productCategorys, http.StatusOK)
}

func getProductCategory(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	productCategory, err := s.ProductCategory().Get(id)
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve productCategory", http.StatusInternalServerError)
		return
	}

	respond(w, productCategory, http.StatusOK)
}

func createProductCategory(w http.ResponseWriter, r *http.Request) {
	var productCategory model.ProductCategory

	err := readBytes(w, r, &productCategory)
	if err != nil {
		respondMsg(w, "Error: Failed to create productCategory", http.StatusBadRequest)
		return
	}

	err = validation.ValidateProductCategory(productCategory)
	if err != nil {
		respondMsg(w, "Error: Invalid productCategory data", http.StatusBadRequest)
		return
	}

	err = s.ProductCategory().Create(&productCategory)
	if err != nil {
		respondMsg(w, "Error: Failed to create productCategory", http.StatusInternalServerError)
		return
	}

	respond(w, productCategory, http.StatusOK)
}

func updateProductCategory(w http.ResponseWriter, r *http.Request) {
	var productCategory model.ProductCategory

	err := readBytes(w, r, &productCategory)
	if err != nil {
		return
	}

	err = validation.ValidateProductCategory(productCategory)
	if err != nil {
		respondMsg(w, "Error: Invalid productCategory data", http.StatusBadRequest)
		return
	}

	err = s.ProductCategory().Update(&productCategory)
	if err != nil {
		respondMsg(w, "Error: Failed to update productCategory", http.StatusInternalServerError)
		return
	}

	respond(w, productCategory, http.StatusOK)
}

func deleteProductCategory(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	err = s.ProductCategory().Delete(id)
	if err != nil {
		respondMsg(w, "Error: Failed to delete productCategory", http.StatusInternalServerError)
		return
	}

	respondEmpty(w, http.StatusOK)
}
