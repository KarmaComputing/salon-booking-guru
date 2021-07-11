package handler

import (
	"net/http"
	"salon-booking-guru/store/model"
)

func accountRoutes() {
	router.HandleFunc("/account", getAllAccount).Methods("GET")
	router.HandleFunc("/account/{id}", getAccount).Methods("GET")
	router.HandleFunc("/account", createAccount).Methods("POST")
	router.HandleFunc("/account", updateAccount).Methods("PUT")
	router.HandleFunc("/account/{id}", deleteAccount).Methods("DELETE")
}

func getAllAccount(w http.ResponseWriter, r *http.Request) {
	accounts, err := s.Account().GetAll()
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve accounts", http.StatusInternalServerError)
		return
	}

	respond(w, accounts, http.StatusOK)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	account, err := s.Account().Get(id)
	if err != nil {
		respondMsg(w, "Error: Failed to retrieve account", http.StatusInternalServerError)
		return
	}

	respond(w, account, http.StatusOK)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account

	err := readBytes(w, r, &account)
	if err != nil {
		return
	}

	err = s.Account().Create(&account)
	if err != nil {
		respondMsg(w, "Error: Failed to create account", http.StatusInternalServerError)
		return
	}

	respond(w, account, http.StatusOK)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account

	err := readBytes(w, r, &account)
	if err != nil {
		return
	}

	err = s.Account().Update(&account)
	if err != nil {
		respondMsg(w, "Error: Failed to update account", http.StatusInternalServerError)
		return
	}

	respond(w, account, http.StatusOK)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r, "id")
	if err != nil {
		return
	}

	err = s.Account().Delete(id)
	if err != nil {
		respondMsg(w, "Error: Failed to delete account", http.StatusInternalServerError)
		return
	}

	respondEmpty(w, http.StatusOK)
}
