package main

import (
	"log"
	"strings"

	"github.com/davidpolme/mutant-detector/script-logic/utils"
)

func main() {
	dnaRequest := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAGGG", "CCCCTA", "TCACTG"}

	//Convert slices to string
	dnaString := strings.Join(dnaRequest, "")
	//Separate string into slices of strings
	dna := utils.StringToSlice(dnaString, 6)
	//Convert slices to matrix
	dnaMatrix := utils.SliceToMatrix(dna)
	//Check if there is a pattern
	dnaResponse := checkIfMutant(dnaMatrix)
	log.Println("[result]:", dnaResponse)
}

//checkIfMutant is used to check if there is an anomaly pattern in the dna sequence
//Inputs: Matrix of DNA sequence
//Outputs: true if there is an anomaly pattern in the dna sequence
func checkIfMutant(dna [][]string) bool {
	count := 0
	count += checkHorizontal(dna)
	if count > 1 {
		return true
	}
	count += checkVertical(dna)
	if count > 1 {
		return true
	}
	count += checkDiagonalPositive(dna)
	if count > 1 {
		return true
	}
	count += checkDiagonalNegative(dna)
	return count > 1
}

func checkHorizontal(dna [][]string) int {
	//fmt.Println("[Matrix]", dna)
	count := 0

	for i := 0; i < len(dna); i++ {
		//search hints for possible patterns
		if dna[i][0] == dna[i][2] {
			if dna[i][0] == dna[i][1] && dna[i][1] == dna[i][3] {
				count++
				continue
			}
		}
		if dna[i][3] == dna[i][4] {
			if dna[i][2] == dna[i][3] && dna[i][1] == dna[i][3] {
				count++
				continue
			}
		}
	}
	log.Println("[count]", count)
	return count
}

func checkVertical(dna [][]string) int {
	dnaMatrix := utils.TransposeMatrix(dna)
	count := checkHorizontal(dnaMatrix)
	return count
}

func checkDiagonalNegative(dna [][]string) int {
	count := 0
	return count
}

func checkDiagonalPositive(dna [][]string) int {
	dnaMatrix := utils.ReverseMatrix(dna)
	count := checkDiagonalNegative(dnaMatrix)
	return count
}
