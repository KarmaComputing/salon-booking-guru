package main

import (
	"log"
	"net/http"
	"os"
	"salon-booking-guru/handler"
	"salon-booking-guru/store"
	"salon-booking-guru/store/psqlstore"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var s store.Store
	var err error

	// Initialise instance of store.Store
	s, err = psqlstore.Open()
	if err != nil {
		log.Fatal("Fatal: Can't start the server without a store.")
		return
	}

	router := mux.NewRouter()
	router.Use(handler.CommonMiddleware)

	handler.InitRouter(router, s)

	log.Println("Listening on :8085")
	log.Println(http.ListenAndServe(":8085", ghandlers.CORS(
		ghandlers.AllowedHeaders([]string{
			"Access-Control-Allow-Origin",
			"X-Requested-With",
			"Content-Type",
			"Authorization",
			"Headers",
			"ResponseType",
		}),
		ghandlers.AllowedMethods([]string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"HEAD",
			"OPTIONS",
		}),
		ghandlers.AllowedOrigins([]string{
			os.Getenv("ORIGIN_ALLOWED"),
			"*",
		},
		))(router)))
}
