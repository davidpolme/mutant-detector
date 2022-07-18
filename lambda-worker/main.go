package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davidpolme/mutant-detector/lambda-worker/handler"
)

func main() {
	lambda.Start(handler.SQSHandler)
}
