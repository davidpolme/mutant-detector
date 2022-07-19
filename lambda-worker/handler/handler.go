package handler

import (
	"context"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/davidpolme/mutant-detector/lambda-worker/controller"
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

func SQSHandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {

		logger.Info("Recieved SQS event", zap.Any("message", message.Body))

		//check if the sequence is a mutant
		res := controller.CheckMutant(message.Body)

		//update dynamodbattribute
		dna_obj := models.DnaSeq{Id: message.Body, IsMutant: strconv.FormatBool(res), Status: "processed"}

		logger.Info("Updating DynamoDB", zap.Any("dna_obj", dna_obj))

		_, err := db.UpdateDnaSeq(dna_obj)
		if err != nil {
			logger.Error("Error updating DynamoDB", zap.Error(err))
			return err
		}

	}
	return nil
}
