package aws

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/davidpolme/mutant-detector/orchestator-service/config"
)

// SendMessageToSQS send a message to SQS
// Inputs: message
func SendMessageToSQS(msg string) {
	sess, _ := session.NewSession(&aws.Config{
		Region:     aws.String(config.Region),
		MaxRetries: aws.Int(5),
	})

	svc := sqs.New(sess)

	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:            aws.String(msg), // Required
		QueueUrl:               aws.String(config.QueueUrl), // Required
		MessageGroupId:         aws.String("1"),
		MessageDeduplicationId: aws.String("secret"), // Required    
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("[Send message] \n%v \n\n", send_resp)
}
