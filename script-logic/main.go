package main

import (
	"log"
	"strings"

	"github.com/davidpolme/mutant-detector/script-logic/controllers"
	"github.com/davidpolme/mutant-detector/script-logic/utils"
)

func main() {
	dnaRequest := []string{"TCACTG", "CCACTA", "AACACA", "CACACA", "CAGTGC", "AGGGTT"}

	//Convert slices to string
	dnaString := strings.Join(dnaRequest, "")
	//Separate string into slices of strings
	dna := utils.StringToSlice(dnaString, 6)
	//Convert slices to matrix
	dnaMatrix := utils.SliceToMatrix(dna)
	//Check if there is a pattern
	dnaResponse := controllers.CheckIfMutant(dnaMatrix)
	log.Println("[result]:", dnaResponse)
}
