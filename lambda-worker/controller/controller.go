package controller

import (
	"log"

	"github.com/davidpolme/mutant-detector/lambda-worker/utils"
)

func CheckMutant(dnaRequest string) bool {
	dna := utils.StringToSlice(dnaRequest, 6)
	//Convert slices to matrix
	dnaMatrix := utils.SliceToMatrix(dna)
	//Check if there is a pattern
	dnaResponse := checkIfMutant(dnaMatrix)
	log.Println("[result]:", dnaResponse)
	return dnaResponse
}

//checkIfMutant is used to check if there is an anomaly pattern in the dna sequence
//Inputs: Matrix of DNA sequence
//Outputs: true if there is an anomaly pattern in the dna sequence
func checkIfMutant(dna [][]string) bool {
	count := 0
	count = checkHorizontal(dna, count)
	//log.Println("[count horizontal]", count)
	if count > 1 {
		return true
	}
	count = checkVertical(dna, count)
	//log.Println("[count vertical]", count)

	if count > 1 {
		return true
	}
	count = checkDiagonalNegative(dna, count)

	if count > 1 {
		return true
	}

	count = checkDiagonalPositive(dna, count)
	return count > 1
}

func checkHorizontal(dna [][]string, count int) int {

	for i := 0; i < len(dna); i++ {
		//Si en esta secuencia count es mayor que 1 se retorna el valor de count
		if count > 1 {
			break
		}
		//search hints for possible patterns
		if dna[i][0] == dna[i][3] {
			if dna[i][0] == dna[i][1] && dna[i][1] == dna[i][2] {
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
	return count
}

func checkVertical(dna [][]string, count int) int {
	dnaMatrix := utils.TransposeMatrix(dna)
	count = checkHorizontal(dnaMatrix, count)
	//log.Println("[count Vertical]", count)
	return count
}

func checkDiagonalNegative(dna [][]string, count int) int {
	//1 step: Check main diagonal
	count = utils.CheckMainDiagonal(dna, count)

	if count > 1 {
		return count
	}

	//2 step, check adjacent diagonals
	count = utils.CheckAdjacentDiagonals(dna, count)
	return count
}

func checkDiagonalPositive(dna [][]string, count int) int {
	dnaMatrix := utils.ReverseMatrix(dna)
	count = checkDiagonalNegative(dnaMatrix, count)
	//log.Println("[count d-positive]: ", count)
	return count
}
