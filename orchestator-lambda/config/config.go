package config

var (
	// Endpoint
	URL              = "https://9pto8im9dl.execute-api.us-east-1.amazonaws.com"
	EndpointDetector = "/detector"
	Region           = "us-east-1"
	DynamoTable      = "dnaSequences"
	QueueUrl         = "https://sqs.us-east-1.amazonaws.com/146521158052/dna_anomaly_queue.fifo"
)
