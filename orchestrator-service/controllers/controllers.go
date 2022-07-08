package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davidpolme/mutant-detector/orchestator-service/aws"
	"github.com/davidpolme/mutant-detector/orchestator-service/models"
	"github.com/davidpolme/mutant-detector/orchestator-service/utils"
)

// ValidateMutant recieve a dna sequence then validate if its a mutant or not.
// Inputs:
//     w: http.ResponseWriter - The response writer.
//     r: *http.Request - The HTTP request.
// Output:
//     If success, a SendMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to SendMessage.
func ValidateMutant(w http.ResponseWriter, r *http.Request) {
	var dnaStruct models.DnaSeq

	//recieve data and send post request
	err := json.NewDecoder(r.Body).Decode(&dnaStruct)
	if err != nil {
		log.Printf("Error decoding JSON: %v  \n", err)
	}

	//Set id equal to DNA sequence
	dnaStruct.Id = strings.Join(dnaStruct.Dna, "")

	err = utils.CheckADNStruct(dnaStruct.Dna, dnaStruct.Id)
	if err != nil {
		log.Printf("Error validating DNA: %v  \n", err)
	}

	//Check if exist in DB
	mutantStatus := checkIfMutantFromDb(dnaStruct.Id)
	if mutantStatus != "Undetermined" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"isMutant":  "` + mutantStatus + `"}`))
		return
	}

	//Encode data to json
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(dnaStruct)
	if err != nil {
		log.Fatal(err)
	}
	//save to db
	err = InsertIntoDb(&buf)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error":"` + err.Error() + `"}`))
		return
	}
	/*
		// TODO:
		ismutant := getFromDb(dnaStruct.Id)
			if ismutant == "true" {
			}


			if isMutant {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(id))
			}
	*/
	aws.SendMessageToSQS(dnaStruct.Id)
	log.Println(dnaStruct.Id)

	//data := getFromDb(id)

	//Check if exist in DB
	mutantStatus = checkIfMutantFromDb(dnaStruct.Id)
	if mutantStatus != "Undetermined" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"isMutant":  "` + mutantStatus + `"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"message":  "Se está procesando la peticion, por favor inténtelo de nuevo"}`))
}

func InsertIntoDb(buf *bytes.Buffer) error {
	//send data to post
	response, err := http.Post("http://localhost:8081/db", "application/json", buf)
	if err != nil {
		log.Printf("Error sending POST request: %v  \n", err)
		return err
	}
	defer response.Body.Close()

	//decode response
	var msgStruct models.Message
	err = json.NewDecoder(response.Body).Decode(&msgStruct)
	if err != nil {
		log.Printf("Error decoding JSON: %v \n", err)
		return err
	}
	return nil
}

//

func checkIfMutantFromDb(s string) (string) {
	//Get data

	//create json object
	newjson := `{"id": "` + s + `"}`

	//send data to post
	response, err := http.Post("http://localhost:8081/db-exist", "application/json", bytes.NewBuffer([]byte(newjson)))
	if err != nil {
		log.Printf("Error getting data from DB: %v  \n", err)
		return "Undetermined"
	}

	//decode response
	var dnaStruct models.DnaSeq

	err = json.NewDecoder(response.Body).Decode(&dnaStruct)
	if err != nil {
		log.Printf("Error decoding JSON: %v \n", err)
		return "Undetermined"
	}

	if dnaStruct.IsMutant == "" {
		return "Undetermined"
	}

	fmt.Println("Response: ", dnaStruct)
	return dnaStruct.IsMutant
}

func SendHello(w http.ResponseWriter, r *http.Request) {
	var helloStruct models.Message

	//print body message
	err := json.NewDecoder(r.Body).Decode(&helloStruct)
	if err != nil {
		fmt.Println(err)
		return
	}
	//print body message
	fmt.Println("Mensaje Original: ", helloStruct.Message)

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(helloStruct)
	if err != nil {
		log.Fatal(err)
	}
	//send data to post
	response, err := http.Post("http://localhost:8081/hello", "application/json", &buf)

	if err != nil {
		log.Fatal(err)
	}
	//
	//decode response
	var helloResponse models.Message
	err = json.NewDecoder(response.Body).Decode(&helloResponse)
	if err != nil {
		log.Fatal(err)
	}
	//print body message
	fmt.Println("Mensaje Respuesta: ", helloResponse.Message)
}
