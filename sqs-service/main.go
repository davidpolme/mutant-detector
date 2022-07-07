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
		MessageBody:            aws.String("message body by dpmsnotes Hey"), // Required
		QueueUrl:               aws.String(conf.QueueUrl),                 // Required
		MessageGroupId:         aws.String("20"),
		MessageDeduplicationId: aws.String("ddd"), // Required             // (optional) 傳進去的 message 延遲 n 秒才會被取出, 0 ~ 900s (15 minutes)
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[Send message] \n%v \n\n", send_resp)
}
