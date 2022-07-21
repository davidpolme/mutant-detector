package db

import (
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/config"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/models"
)

var Dynamo *dynamodb.DynamoDB

func init() {
	Dynamo = ConnectDynamo()
}

// ConnectDynamo creates a new dynamo session and returns a pointer to the dynamo client
func ConnectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
}
func SendToDynamoDB(requestBody models.Request, isMutant bool) error {
	dna_id := strings.Join(requestBody.DNA, "")
	_, err := Dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: &config.DynamoTable,
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(dna_id),
			},
			"IsMutant": {
				S: aws.String(strconv.FormatBool(isMutant)),
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
