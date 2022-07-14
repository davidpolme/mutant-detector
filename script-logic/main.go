package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/davidpolme/mutant-detector/script-logic/utils"
)

func main() {
	dnaRequest := []string{"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"}

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
	/*count = checkHorizontal(dna, count)
	if count > 1 {
		return true
	}
	count += checkVertical(dna, count)
	if count > 1 {
		return true
	}*/
	count += checkDiagonalNegative(dna)
	if count > 1 {
		return true
	}
	count += checkDiagonalPositive(dna)
	return count > 1
}

func checkHorizontal(dna [][]string, count int) int {
	fmt.Println("[Matrix]:", dna)

	for i := 0; i < len(dna); i++ {
		//Si en esta secuencia count es mayor que 1 se retorna el valor de count
		if count > 1 {
			break
		}
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

func checkVertical(dna [][]string, count int) int {
	dnaMatrix := utils.TransposeMatrix(dna)
	count = checkHorizontal(dnaMatrix, count)
	return count
}

func checkDiagonalNegative(dna [][]string) int {
	log.Println("[dna]: ", dna)
	count := 0

	//1 step: Check diagonal
	if dna[0][0] == dna[3][3]{
		if dna[1][1] == dna[2][2]{
			count++
		}
	}else if dna[3][3] == dna[4][4]{
		if dna[2][2] == dna[5][5]{
			count++
		}

		adjX1 := 0
		adjX2 := 3
		centerX:=3
		centerY:=2
	//2 step, check adjacent diagonals
	for i := 1; i <= 2; i++ {
		if dna[adjX1][i] == dna[adjX2][i+adjX2]{
			if dna[adjX1+1][i+1] == dna[adjX2-1][i+2]{
				count++
			}else if i == 1{
				if dna[centerX][centerY] == dna[centerX+1][centerY+1]{
					count++
				}
			}	
		}
	}
	//3 step, check middle 
	}
	
	return count
}

func checkDiagonalPositive(dna [][]string) int {
	dnaMatrix := utils.ReverseMatrix(dna)
	count := checkDiagonalNegative(dnaMatrix)
	return count
}
