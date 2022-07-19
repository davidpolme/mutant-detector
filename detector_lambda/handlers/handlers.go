package handlers

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/davidpolme/mutant-detector/detector_lambda/models"
	"github.com/davidpolme/mutant-detector/detector_lambda/utils"
)

func MyHandler(request events.APIGatewayProxyRequest) (models.Response, error) {
	// Unmarshal the request body into the models.Request struct
	var requestBody models.Request
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		log.Println("Error unmarshalling request body:", err)
		return models.Response{}, err
	}

	isMutant := utils.IsMutant(requestBody.DNA)

	return models.Response{
		IsMutant: isMutant,
	}, nil
}
