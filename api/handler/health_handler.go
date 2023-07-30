package handler

import (
	"net/http"
)

func healthRoutes() {
	router.HandleFunc("/health", getHealth).Methods("GET")
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	respond(w, '{"status": "OK"}', http.StatusOK)
}

