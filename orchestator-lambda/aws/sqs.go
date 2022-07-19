package aws

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

	var sqsMessagestruct models.SQSMessage
	sqsMessagestruct.Request = request
	sqsMessagestruct.Response = b

	//Marshal request
	sqsMessage, err := json.Marshal(sqsMessagestruct)
	if err != nil {
		return err
	}
	//convert request to string
	requestString := string(sqsMessage)

	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:            aws.String(requestString),   // Required
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
