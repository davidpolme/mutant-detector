package models

type DnaSeq struct {
	DnaId       string   `json:"id"`
	IsMutant string   `json:"isMutant"`
	Status   string   `json:"status"`
}
