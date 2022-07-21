package sqs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/config"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/models"
)

func SendToSQS(request models.Request, b bool) error {
	sess, _ := session.NewSession(&aws.Config{
		Region:     aws.String(config.Region),
		MaxRetries: aws.Int(5),
	})

	svc := sqs.New(sess)

	//Create message
	message := models.SQSMessage{
		IsMutant: b,
		DNA:      request.DNA,
	}
	
	//Marshal message
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return err
	}

	//convert request to string
	requestString := string(jsonMessage)

	//print log
	log.Println("Sending message to SQS:", requestString)

	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:            aws.String(requestString),      // Required
		QueueUrl:               aws.String(config.QueueUrl), // Required
		MessageGroupId:         aws.String("1"),
		MessageDeduplicationId: aws.String("secret"), // Required
	}
	
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("[Send message] \n%v \n\n", send_resp)

	return nil
}
