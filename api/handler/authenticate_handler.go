package handler

import (
	"net/http"
	"salon-booking-guru/store/model"
)

func authenticateRoutes() {
	v1.HandleFunc("/authenticate", authenticateCredentials).Methods("POST")
}

func authenticateCredentials(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials

	err := readBytes(w, r, &credentials)
	if err != nil {
		respondMsg(w, "Error: Failed to authenticate credentials", http.StatusBadRequest)
		return
	}

	authenticateResponse, err := s.Authenticate().AuthenticateCredentials(credentials)
	if err != nil {
		respondMsg(w, "Error: Failed to authenticate credentials", http.StatusBadRequest)
		return
	}

	respond(w, authenticateResponse, http.StatusOK)
}
