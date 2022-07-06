package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/davidpolme/mutant-detector/orchestator-service/models"
)

func ValidateMutant(w http.ResponseWriter, r *http.Request) {
	response, err := http.Post("http://localhost:8081/db", "application/json", r.Body)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	
	if err != nil {
		log.Fatal(err)
	}
	
	var responseObject models.Hello
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + responseObject.Message + `}`))
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8081/hello")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject models.Hello
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Message)

	newResponse := responseObject.Message + "; Received from Orchestrator Service"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + newResponse + `}`))
}
