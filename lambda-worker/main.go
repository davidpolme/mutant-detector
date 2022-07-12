package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/davidpolme/mutant-detector/lambda-worker/config"
	"github.com/davidpolme/mutant-detector/lambda-worker/models"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	l, _ := zap.NewProduction()
	logger = l
	defer logger.Sync() // flushes buffer, if any

}

func dynamoHandler(id string) error {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(session)

	result, err := dynamodbattribute.MarshalMap(id)
	if err != nil {
		fmt.Println("Failed to marshall request")
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      "message",
		TableName: aws.String(config.DynamoTable),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		logger.Info("Failed to write to db")
		return err
	}
	return nil
}

func MyHandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {

		logger.Info("Recieved SQS event", zap.Any("message", message.Body))


		// Do work to process the message
		// ...

		//update dynamodbattribute

	}
	return nil
}

func main() {
	lambda.Start(MyHandler)
}
