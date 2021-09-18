package handler

import (
	"net/http"
)

func roleRoutes() {
	// GET
	v1.HandleFunc(
		"/role",
		authorize(getAllRole, "canReadRole"),
	).Methods("GET")
}

func getAllRole(w http.ResponseWriter, r *http.Request) {
	roles, err := s.Role().GetAll()
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve roles", http.StatusInternalServerError)
		return
	}

	respond(w, roles, http.StatusOK)
}
