package handler

import (
	"net/http"
	"salon-booking-guru/store/model"
	"salon-booking-guru/validation"
)

func availabilityRoutes() {
	// GET
	v1.HandleFunc(
		"/availability",
		authorize(getAllAvailability, "canReadAvailability"),
	).Methods("GET")
	v1.HandleFunc(
		"/availability/{id}",
		authorize(getAvailability, "canReadAvailability"),
	).Methods("GET")

	// POST
	v1.HandleFunc(
		"/availability",
		authorize(createAvailability, "canCreateAvailability"),
	).Methods("POST")

	// PUT
	v1.HandleFunc(
		"/availability",
		authorize(updateAvailability, "canUpdateAvailability"),
	).Methods("PUT")

	// DELETE
	v1.HandleFunc(
		"/availability/{id}",
		authorize(deleteAvailability, "canDeleteAvailability"),
	).Methods("DELETE")
}

func getAllAvailability(w http.ResponseWriter, r *http.Request) {
	availabilitys, err := s.Availability().GetAll()
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve availabilitys", http.StatusInternalServerError)
		return
	}

	respond(w, availabilitys, http.StatusOK)
}

func getAvailability(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	availability, err := s.Availability().Get(id)
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve availability", http.StatusInternalServerError)
		return
	}

	respond(w, availability, http.StatusOK)
}

func createAvailability(w http.ResponseWriter, r *http.Request) {
	var availability model.Availability

	err := readBytes(w, r, &availability)
	if err != nil {
		respondMsg(w, "Error: Failed to create availability", http.StatusBadRequest)
		return
	}

	err = validation.ValidateAvailability(availability)
	if err != nil {
		respondMsg(w, "Error: Invalid availability data", http.StatusBadRequest)
		return
	}

	err = s.Availability().Create(&availability)
	if err != nil {
		respondMsg(w, "Error: Failed to create availability", http.StatusInternalServerError)
		return
	}

	respond(w, availability, http.StatusOK)
}

func updateAvailability(w http.ResponseWriter, r *http.Request) {
	var availability model.Availability

	err := readBytes(w, r, &availability)
	if err != nil {
		return
	}

	err = validation.ValidateAvailability(availability)
	if err != nil {
		respondMsg(w, "Error: Invalid availability data", http.StatusBadRequest)
		return
	}

	err = s.Availability().Update(&availability)
	if err != nil {
		respondMsg(w, "Error: Failed to update availability", http.StatusInternalServerError)
		return
	}

	respond(w, availability, http.StatusOK)
}

func deleteAvailability(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	err = s.Availability().Delete(id)
	if err != nil {
		respondMsg(w, "Error: Failed to delete availability", http.StatusInternalServerError)
		return
	}

	respondEmpty(w, http.StatusOK)
}
