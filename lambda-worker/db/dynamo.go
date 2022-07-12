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

func ConnectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  	})))
}

func UpdateDnaSeq(dnaseq models.Request) (*dynamodb.UpdateItemOutput, error) {
	return Dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(config.DynamoTable),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(dnaseq.Id),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":dna": {
				S: aws.String(dnaseq.Id),
			},
		},
		UpdateExpression: aws.String("SET dna = :dna"),
	})
}
