package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

	//save to db
	id, err := InsertIntoDb(w, r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error":"` + err.Error() + `"}`))
		return
	}
	log.Println(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Transaccion exitosa"}`))
}

//
func InsertIntoDb(w http.ResponseWriter, r *http.Request) (string, error) {
	var dnaStruct models.DnaSeq

	//recieve data and send post request
	err := json.NewDecoder(r.Body).Decode(&dnaStruct)
	if err != nil {
		log.Printf("Error decoding JSON: %v  \n", err)
		return "", err
	}
	//Set id equal to DNA sequence
	dnaStruct.Id = strings.Join(dnaStruct.Dna, "")

	err = utils.CheckADNStruct(dnaStruct.Dna, dnaStruct.Id)

	if err != nil {
		log.Printf("Error validating DNA: %v  \n", err)
		return "", err
	}
	//Encode data to json
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(dnaStruct)
	if err != nil {
		log.Printf("Error encoding JSON: %v  \n", err)
		return "", err
	}

	//send data to post
	response, err := http.Post("http://localhost:8081/db", "application/json", &buf)
	if err != nil {
		log.Printf("Error sending POST request: %v  \n", err)
		return "", err
	}
	defer response.Body.Close()

	//decode response
	var msgStruct models.Hello
	err = json.NewDecoder(response.Body).Decode(&msgStruct)
	if err != nil {
		log.Printf("Error decoding JSON: %v \n", err)
		return "", err
	}
	return dnaStruct.Id, nil
}

func SendHello(w http.ResponseWriter, r *http.Request) {
	var helloStruct models.Hello

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
	var helloResponse models.Hello
	err = json.NewDecoder(response.Body).Decode(&helloResponse)
	if err != nil {
		log.Fatal(err)
	}
	//print body message
	fmt.Println("Mensaje Respuesta: ", helloResponse.Message)
}
