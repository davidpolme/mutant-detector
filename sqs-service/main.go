package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	QueueUrl    = "https://sqs.us-east-1.amazonaws.com/146521158052/dna_anomaly_queue.fifo"
	Region      = "us-east-1"
	CredPath    = "/home/davidpolme/.aws/credentials"
	CredProfile = "aws-cred-profile"
)

func main() {
	sess, _ := session.NewSession(&aws.Config{
		Region:     aws.String(Region),
		MaxRetries: aws.Int(5),
	})

	svc := sqs.New(sess)

	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:            aws.String("message body by dpmsnotes Hey"), // Required
		QueueUrl:               aws.String(QueueUrl),                    // Required
		MessageGroupId:         aws.String("20"),
		MessageDeduplicationId: aws.String("ddd"), // Required             // (optional) 傳進去的 message 延遲 n 秒才會被取出, 0 ~ 900s (15 minutes)
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[Send message] \n%v \n\n", send_resp)

}
