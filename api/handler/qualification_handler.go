package handler

import (
	"net/http"
	"salon-booking-guru/store/model"
	"salon-booking-guru/validation"
)

func qualificationRoutes() {
	// GET
	v1.HandleFunc(
		"/qualification",
		authorize(getAllQualification, "canReadQualification"),
	).Methods("GET")
	v1.HandleFunc(
		"/qualification/{id}",
		authorize(getQualification, "canReadQualification"),
	).Methods("GET")

	// POST
	v1.HandleFunc(
		"/qualification",
		authorize(createQualification, "canCreateQualification"),
	).Methods("POST")

	// PUT
	v1.HandleFunc(
		"/qualification",
		authorize(updateQualification, "canUpdateQualification"),
	).Methods("PUT")

	// DELETE
	v1.HandleFunc(
		"/qualification/{id}",
		authorize(deleteQualification, "canDeleteQualification"),
	).Methods("DELETE")
}

func getAllQualification(w http.ResponseWriter, r *http.Request) {
	qualifications, err := s.Qualification().GetAll()
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve qualifications", http.StatusInternalServerError)
		return
	}

	respond(w, qualifications, http.StatusOK)
}

func getQualification(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	qualification, err := s.Qualification().Get(id)
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve qualification", http.StatusInternalServerError)
		return
	}

	respond(w, qualification, http.StatusOK)
}

func createQualification(w http.ResponseWriter, r *http.Request) {
	var qualification model.Qualification

	err := readBytes(w, r, &qualification)
	if err != nil {
		respondMsg(w, "Error: Failed to create qualification", http.StatusBadRequest)
		return
	}

	err = validation.ValidateQualification(qualification)
	if err != nil {
		respondMsg(w, "Error: Invalid qualification data", http.StatusBadRequest)
		return
	}

	err = s.Qualification().Create(&qualification)
	if err != nil {
		respondMsg(w, "Error: Failed to create qualification", http.StatusInternalServerError)
		return
	}

	respond(w, qualification, http.StatusOK)
}

func updateQualification(w http.ResponseWriter, r *http.Request) {
	var qualification model.Qualification

	err := readBytes(w, r, &qualification)
	if err != nil {
		return
	}

	err = validation.ValidateQualification(qualification)
	if err != nil {
		respondMsg(w, "Error: Invalid qualification data", http.StatusBadRequest)
		return
	}

	err = s.Qualification().Update(&qualification)
	if err != nil {
		respondMsg(w, "Error: Failed to update qualification", http.StatusInternalServerError)
		return
	}

	respond(w, qualification, http.StatusOK)
}

func deleteQualification(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	err = s.Qualification().Delete(id)
	if err != nil {
		respondMsg(w, "Error: Failed to delete qualification", http.StatusInternalServerError)
		return
	}

	respondEmpty(w, http.StatusOK)
}
