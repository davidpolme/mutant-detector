package main

import(
	"github.com/davidpolme/mutant-detector/db"
	"github.com/davidpolme/mutant-detector/models"
)
var(
	Id = 1
)

func init() {
  db.Dynamo = db.ConnectDynamo()
}

func main() {
	err := db.CreateTable()



  // make new dna sequence
  dnaseq1 := models.DnaSeq {
    Id: Id,
    Sequence: {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"},
  }
	
}