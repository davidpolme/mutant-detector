package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davidpolme/mutant-detector/orchestator-lambda/config"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/models"
)

func ValidateRequest(request models.Request) error {
	isMatrix := IsMatrix(request.DNA)
	if !isMatrix {
		return errors.New("invalid matrix")
	}

	validChars := ContainsValidChars(request.DNA)
	if !validChars {
		return errors.New("invalid matrix")
	}
	return nil
}

func IsMutant(request models.Request) (models.Response, error) {
	//Encode the data
	postBody, _ := json.Marshal(map[string][]string{
		"dna": request.DNA,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Create the request
	resp, err := http.Post(config.URL+config.EndpointDetector, "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	//Handle Error
	if err != nil {
		log.Fatalln(err)
	}
	//Unmarshal the response body
	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}
	//Return the response
	return response, nil
}
