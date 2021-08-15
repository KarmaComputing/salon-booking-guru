package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"salon-booking-guru/store"
	"strconv"

	"github.com/gorilla/mux"
)

var router *mux.Router
var s store.Store
var v1 *mux.Router

func InitRouter(r *mux.Router, mainStore store.Store) {
	s = mainStore
	router = r
	v1 = router.PathPrefix("/v1").Subrouter()
	accountRoutes()
	authenticateRoutes()
	availabilityRoutes()
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func respond(w http.ResponseWriter, model interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model)
}

func respondMsg(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(
		struct {
			Message string `json:"message"`
		}{
			message,
		},
	)
}

func respondEmpty(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func readBytes(w http.ResponseWriter, r *http.Request, model interface{}) error {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		respondMsg(w, "Error: Failed to read request body", http.StatusInternalServerError)
		return err
	}
	json.Unmarshal(bodyBytes, model)

	return nil
}

func getId(w http.ResponseWriter, r *http.Request, idName string) (int, error) {
	idStr := mux.Vars(r)[idName]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorMsg := fmt.Sprintf("Error: {%s} is not an integer", idName)
		log.Println(errorMsg)
		log.Println(err)
		respondMsg(w, errorMsg, http.StatusInternalServerError)
		return 0, err
	}

	return id, nil
}
