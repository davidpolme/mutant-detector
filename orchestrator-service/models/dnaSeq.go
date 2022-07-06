package models

type DnaSeq struct {
	Id 		string 		`json:"id"`
	Dna		[]string 	`json:"dna"`
}