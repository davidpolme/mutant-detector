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
	_, err := Dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: &config.TableName,
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

func GetDnaSeq(id string) (models.DnaSeq, error) {
	var dnaseq models.DnaSeq
	result, err := Dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: &config.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return dnaseq, err
	}
	if result.Item == nil {
		return dnaseq, err
	}
	dnaseq.Id = *result.Item["Id"].S
	dnaseq.IsMutant = *result.Item["IsMutant"].S
	dnaseq.Status = *result.Item["Status"].S
	return dnaseq, nil
}

func UpdateDnaSeq(dnaseq models.DnaSeq) (*dynamodb.UpdateItemOutput, error) {
	d, err := Dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: &config.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(dnaseq.Id),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#Dna": aws.String("Dna"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":dna": {
				SS: aws.StringSlice(dnaseq.Dna),
			},
		},
		UpdateExpression: aws.String("SET #Dna = :dna"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	})
	return d, err
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
