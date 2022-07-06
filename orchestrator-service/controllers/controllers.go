package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/davidpolme/mutant-detector/orchestator-service/models"
)

func InsertIntoDB(w http.ResponseWriter, r *http.Request) string {
	response, err := http.Post("http://localhost:8081/db", "application/json", r.Body)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject models.DnaSeq
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Dna)

	return  responseObject.Id
}


// ValidateMutant recieve a dna sequence then validate if its a mutant or not.
// Inputs:
//     w: http.ResponseWriter - The response writer.
//     r: *http.Request - The HTTP request.
// Output:
//     If success, a SendMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to SendMessage.
func ValidateMutant(w http.ResponseWriter, r *http.Request) {
	var dnaStruct models.DnaSeq
	err := json.NewDecoder(r.Body).Decode(&dnaStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(dnaStruct.Dna) == 0 {
		http.Error(w, "Dna is empty", http.StatusBadRequest)
		return
	}
	dnaStruct.Id = strings.Join(dnaStruct.Dna, "")
	json.NewEncoder(w).Encode(dnaStruct)

	response := InsertIntoDB(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + response + `}`))
}
