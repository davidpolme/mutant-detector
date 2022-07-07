package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidpolme/mutant-detector/orchestator-service/config"
	"github.com/davidpolme/mutant-detector/orchestator-service/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Setting port and set enable server
func Handlers() {
	//create router
	router := mux.NewRouter()

	//create route for db
	router.HandleFunc("/mutant", controllers.ValidateMutant).Methods("POST")
	router.HandleFunc("/hello", controllers.SendHello).Methods("POST")

	//set port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = config.PORT
	}
	//allow cors
	fmt.Println("Server started on port " + PORT)
	handler := cors.AllowAll().Handler(router)

	if err := http.ListenAndServe(":"+PORT, handler); err != nil {
		log.Fatal(err)
	}
}
