package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/davidpolme/mutant-detector/sqs-service/conf"
)

func main() {
	sess, _ := session.NewSession(&aws.Config{
		Region:     aws.String(conf.Region),
		MaxRetries: aws.Int(5),
	})

	svc := sqs.New(sess)

	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:            aws.String("message body by dpmsnotes Hey"), 
		QueueUrl:               aws.String(conf.QueueUrl),                
		MessageGroupId:         aws.String("20"),
		MessageDeduplicationId: aws.String("ddd"),        
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[Send message] \n%v \n\n", send_resp)
}
