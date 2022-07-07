package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davidpolme/mutant-detector/orchestator-service/models"
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
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(&dnaStruct)
	//send data to post
	http.Post("http://localhost:8081/mutant", "application/json", r.Body)
	

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
	fmt.Println("Mensaje Original: ",helloStruct.Message)

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(helloStruct)
	if err != nil {
		log.Fatal(err)
	}
	//send data to post
	response, err :=http.Post("http://localhost:8081/hello", "application/json", &buf)

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
	fmt.Println("Mensaje Respuesta: ",helloResponse.Message)
}
