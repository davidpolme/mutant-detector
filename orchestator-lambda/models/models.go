package models

type Request struct {
	DNA []string `json:"dna"`
}

type Response struct {
	IsMutant bool `json:"isMutant"`
}

type SQSMessage struct {
	Request  Request `json:"request"`
	Response bool    `json:"isMutant"`
}
