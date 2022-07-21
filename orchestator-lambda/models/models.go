package models

type Request struct {
	DNA []string `json:"dna"`
}

type Response struct {
	IsMutant bool `json:"isMutant"`
}

type SQSMessage struct {
	DNA      []string `json:"dna"`
	IsMutant bool     `json:"isMutant"`
}
