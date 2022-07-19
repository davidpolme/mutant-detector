package handlers

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/aws"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/models"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/utils"
)

func MyHandler(request events.APIGatewayProxyRequest) (models.Response, error) {
	var requestBody models.Request
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		log.Println("Error unmarshalling request body:", err)
		return models.Response{}, err
	}
	//Check if the request is valid
	err = utils.ValidateRequest(requestBody)
	//Handle Error
	if err != nil {
		return models.Response{}, err
	}

	//Detect if mutant
	resp, err := utils.IsMutant(requestBody)
	//Handle Error
	if err != nil {
		return models.Response{}, err
	}

	//Send Data to SQS Queue to be sent to  database
	err = aws.SendToSQS(requestBody, resp.IsMutant)
	//Handle Error
	if err != nil {
		return models.Response{}, err
	}

	//Return the response
	return resp, nil
}
