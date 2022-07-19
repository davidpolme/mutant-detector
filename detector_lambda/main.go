package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davidpolme/mutant-detector/detector_lambda/handlers"
)

func main() {
	lambda.Start(handlers.MyHandler)
}
