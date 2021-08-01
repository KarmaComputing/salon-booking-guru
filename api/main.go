package main

import (
	"log"
	"net/http"
	"os"
	"salon-booking-guru/handler"
	"salon-booking-guru/store"
	"salon-booking-guru/store/psqlstore"
	"salon-booking-guru/validation"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var s store.Store

	s = psqlstore.Open()

	router := mux.NewRouter()
	router.Use(handler.CommonMiddleware)

	handler.InitRouter(router, s)
	validation.Init(s)

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
