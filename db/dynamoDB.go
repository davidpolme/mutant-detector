package db

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davidpolme/mutant-detector/config"
	"github.com/davidpolme/mutant-detector/models"
)

var Dynamo *dynamodb.DynamoDB

func ConnectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: &config.RegionName,
	})))
}

// CreateTable creates a table
func CreateTable() error {
	_, err := Dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   &config.TableName,
	})

	return err
}

// PutItem inserts the struct Person
func PutItem(dnaseq models.DnaSeq) error {
	_, err := Dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(dnaseq.Id)),
			},
			"Sequence": {
					SS: aws.StringSlice(dnaseq.Sequence),
			},
		},
		TableName: &config.TableName,
	})
	return err
}

