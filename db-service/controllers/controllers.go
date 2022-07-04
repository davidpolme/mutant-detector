package controllers

import (
	"encoding/json"
	"net/http"

	"strings"

	"github.com/davidpolme/mutant-detector/db-service/db"
	"github.com/davidpolme/mutant-detector/db-service/models"
)

func InsertDnaSeq(w http.ResponseWriter, r *http.Request) {
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

	//Valores predeterminados
	dnaStruct.IsMutant = "Undetermined"
	dnaStruct.Status = "Pending"
	dnaStruct.Id = strings.Join(dnaStruct.Dna, "")

	_, err = db.InsertDnaSeq(dnaStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Dna sequence saved in db"}`))
}

func UpdateDnaSeq(w http.ResponseWriter, r *http.Request) {
	/*var dnaStruct models.DnaSeq

	err := json.NewDecoder(r.Body).Decode(&dnaStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(dnaStruct.Dna) == 0 {
		http.Error(w, "Dna is empty", http.StatusBadRequest)
		return
	}

	//err = db.UpdateDnaSeq(dnaStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Dna sequence updated in db"}`))
	*/
}
