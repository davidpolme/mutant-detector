package main

import (
	"fmt"
	"github.com/davidpolme/mutant-detector/db-service/db"
	"github.com/davidpolme/mutant-detector/db-service/handlers"
)

func init() {
	db.Dynamo = db.ConnectDynamo()
	fmt.Println("Connected to DynamoDB")
}

func main() {
	handlers.Handlers()
}
