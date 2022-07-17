package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davidpolme/mutant-detector/lambda-worker/config"
	"github.com/davidpolme/mutant-detector/lambda-worker/models"
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

// InsertDnaSeq inserts the struct Dna into the dynamo table
// Input: DnaSeq struct
// Output: bool, error if any
func UpdateDnaSeq(dnaseq models.DnaSeq) (bool, error) {
	_, err := Dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: &config.DynamoTable,
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(dnaseq.Id),
			},
			"IsMutant": {
				S: aws.String(dnaseq.IsMutant),
			},
			"Status": {
				S: aws.String(dnaseq.Status),
			},
		},
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
