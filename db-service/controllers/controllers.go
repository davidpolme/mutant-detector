package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davidpolme/mutant-detector/db-service/db"
	"github.com/davidpolme/mutant-detector/db-service/models"
)

// GetQueueURL recieve the DNA Sequence to be processed.
// Inputs:
//	w: http.ResponseWriter -> Response to be sent to the client.
//  r: *http.Request -> Request recieved from the client.
// Output:
//     If success, a message is sent to the client.
//     Otherwise, an error is sent to the client.
func GetDnaSeq(w http.ResponseWriter, r *http.Request) {
	var dnaStruct models.DnaSeq

	err := json.NewDecoder(r.Body).Decode(&dnaStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(dnaStruct.Id) == 0 {
		http.Error(w, "Id is empty", http.StatusBadRequest)
		return
	}

	dnaStruct, err = db.GetDnaSeq(dnaStruct.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dnaStruct)
}

func InsertDnaSeq(w http.ResponseWriter, r *http.Request) {
	var dnaStruct models.DnaSeq

	err := json.NewDecoder(r.Body).Decode(&dnaStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if len(dnaStruct.Id) == 0 {
		http.Error(w, "Id is empty", http.StatusBadRequest)
		return
	}
	dnaStruct.Id = strings.Trim(dnaStruct.Id, "\"")
	//Valores predeterminados
	dnaStruct.IsMutant = "Undetermined"
	dnaStruct.Status = "Pending"

	log.Println(dnaStruct.Id)
	_, err = db.InsertDnaSeq(dnaStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Dna sequence inserted in db"}`))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	var helloStruct models.Hello

	err := json.NewDecoder(r.Body).Decode(&helloStruct)

	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helloStruct.Message = helloStruct.Message + " - ;"

	log.Printf("Message: %v", helloStruct.Message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helloStruct)
}

func UpdateDnaSeq(w http.ResponseWriter, r *http.Request) {
	var dnaStruct models.DnaSeq

	err := json.NewDecoder(r.Body).Decode(&dnaStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(dnaStruct.Id) == 0 {
		http.Error(w, "Id is empty", http.StatusBadRequest)
		return
	}

	_, err = db.UpdateDnaSeq(dnaStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Dna sequence updated in db"}`))
}
