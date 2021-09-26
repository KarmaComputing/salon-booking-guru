package handler

import (
	"net/http"
	"salon-booking-guru/store/model"
	"salon-booking-guru/validation"
)

func productRoutes() {
	// GET
	v1.HandleFunc(
		"/product",
		authorize(getAllProduct, "canReadProduct"),
	).Methods("GET")

	v1.HandleFunc(
		"/product/summary",
		authorize(getAllProductSummary, "canReadProduct"),
	).Methods("GET")

	v1.HandleFunc(
		"/product/{id}",
		authorize(getProduct, "canReadProduct"),
	).Methods("GET")

	v1.HandleFunc(
		"/product/{id}/qualification",
		authorize(getAllProductQualificationName, "canReadProduct"),
	).Methods("GET")

	// POST
	v1.HandleFunc(
		"/product",
		authorize(createProduct, "canCreateProduct"),
	).Methods("POST")

	// PUT
	v1.HandleFunc(
		"/product",
		authorize(updateProduct, "canUpdateProduct"),
	).Methods("PUT")

	v1.HandleFunc(
		"/product/{id}/qualification",
		authorize(upsertProductQualification, "canUpdateProduct"),
	).Methods("PUT")

	// DELETE
	v1.HandleFunc(
		"/product/{id}",
		authorize(deleteProduct, "canDeleteProduct"),
	).Methods("DELETE")
}

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := s.Product().GetAll()
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	respond(w, products, http.StatusOK)
}

func getAllProductSummary(w http.ResponseWriter, r *http.Request) {
	productSummaries, err := s.Product().GetAllSummary()
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve product summaries", http.StatusInternalServerError)
		return
	}

	respond(w, productSummaries, http.StatusOK)
}

func getAllProductQualificationName(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	qualificationNames, err := s.Qualification().GetAllNameByProductId(id)
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve account qualification names", http.StatusInternalServerError)
		return
	}

	respond(w, qualificationNames, http.StatusOK)
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
		respondMsg(w, "Error: Failed to create product", http.StatusBadRequest)
		return
	}

	err = validation.ValidateProduct(product)
	if err != nil {
		respondMsg(w, "Error: Invalid product data", http.StatusBadRequest)
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

	err = validation.ValidateProduct(product)
	if err != nil {
		respondMsg(w, "Error: Invalid product data", http.StatusBadRequest)
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

func upsertProductQualification(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	var qualificationIds []int
	err = readBytes(w, r, &qualificationIds)
	if err != nil {
		return
	}

	err = s.Product().UpsertQualification(id, qualificationIds)
	if err != nil {
		respondMsg(w, "Error: Failed to upsert product qualifications", http.StatusInternalServerError)
		return
	}

	qualificationNames, err := s.Qualification().GetAllNameByProductId(id)
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve upserted product qualifications", http.StatusInternalServerError)
		return
	}

	respond(w, qualificationNames, http.StatusOK)
}
