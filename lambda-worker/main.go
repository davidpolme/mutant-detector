package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davidpolme/mutant-detector/lambda-worker/db"
	"github.com/davidpolme/mutant-detector/lambda-worker/models"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	l, _ := zap.NewProduction()
	logger = l
	defer logger.Sync() // flushes buffer, if any

}

func MyHandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {

		logger.Info("Recieved SQS event", zap.Any("message", message.Body))

		// Do work to process the message
		// ...

		//update dynamodbattribute
		dna_obj := models.DnaSeq{Id: message.Body, IsMutant: "true", Status: "processed"}

		logger.Info("Updating DynamoDB", zap.Any("dna_obj", dna_obj))

		_, err := db.UpdateDnaSeq(dna_obj)
		if err != nil {
			logger.Error("Error updating DynamoDB", zap.Error(err))
			return err
		}

	}
	return nil
}

func main() {
	lambda.Start(MyHandler)
}
