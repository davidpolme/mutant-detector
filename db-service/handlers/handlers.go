package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidpolme/mutant-detector/db-service/controllers"
	"github.com/davidpolme/mutant-detector/db-service/db"
	"github.com/davidpolme/mutant-detector/db-service/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func isTableCreated(c chan string) chan string {
	err := db.CreateTable()
	if err != nil {
		c <- "Table already exists"
		return c
	}
	c <- "Table created"
	return c
}

//Setting port and set enable server
func Handlers() {

	//check if table is created
	c := make(chan string)
	go isTableCreated(c)
	fmt.Println(<-c)

	//create router
	router := mux.NewRouter()

	//create route for db
	router.HandleFunc("/db", middlewares.CheckDB(controllers.InsertDnaSeq)).Methods("POST")
	router.HandleFunc("/db-exist", middlewares.CheckDB(controllers.GetDnaSeq)).Methods("POST")

	//set port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	//allow cors
	fmt.Println("Server started on port " + PORT)
	handler := cors.AllowAll().Handler(router)

	if err := http.ListenAndServe(":"+PORT, handler); err != nil {
		log.Fatal(err)
	}
}
