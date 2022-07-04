package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davidpolme/mutant-detector/db-service/config"
	"github.com/davidpolme/mutant-detector/db-service/models"
)

var Dynamo *dynamodb.DynamoDB

func init() {
	Dynamo = ConnectDynamo()
}

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
				AttributeType: aws.String("S"),
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

// InsertDnaSeq inserts the struct Dna
func InsertDnaSeq(dnaseq models.DnaSeq) (bool, error) {
	//TODO: implement cache
	//Consulta en el caché

	//Consulta si existe en la base de datos

	//Si no existe, inserta en la base de datos
	_, err := Dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(dnaseq.Id),
			},
			"Dna": {
				SS: aws.StringSlice(dnaseq.Dna),
			},
			"IsMutant": {
				S: aws.String(dnaseq.IsMutant),
			},
			"Status": {
				S: aws.String(dnaseq.Status),
			},
		},
		TableName: &config.TableName,
	})
	return false, err
}

func getItemById(id string) error {
	_, err := Dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: &config.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	})
	return err
}

func existItemInDB(id string) bool {
	_, err := Dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: &config.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return false
	}
	return true
}
