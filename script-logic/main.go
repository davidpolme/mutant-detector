package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/davidpolme/mutant-detector/script-logic/utils"
)

func main() {
	dnaRequest := []string{"TCACTG", "TTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCAA"}

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

func checkMainDiagonal(dna [][]string, count int) int {
	if count > 1 {
		return count
	}
	if dna[0][0] == dna[3][3] {
		if dna[1][1] == dna[2][2] {
			count++
		}
	} else if dna[3][3] == dna[4][4] {
		if dna[2][2] == dna[5][5] {
			count++
		}
	}
	return count
}

func checkAdjacentDiagonal(dna [][]string, count int) int {
	if count > 1 {
		return count
	}

	adjX1 := 0
	adjX2 := 3
	centerX := 2
	centerY := 3

	//TODO: Separate in two functions
	//TODO: Put those two functions in goroutines
	for i := 1; i <= 2; i++ {
		//Check left diagonals
		if dna[adjX1][i] == dna[adjX2][i+adjX2] {
			if dna[adjX1+1][i+1] == dna[adjX2-1][i+2] {
				count++
			}
		} else if i == 1 {
			//Its necesary to check both of possibilities in the nearest diagonal from  the main diagonal. When i = 1
			if dna[centerX][centerY] == dna[centerX+1][centerY+1] {
				if dna[centerX-1][centerY-1] == dna[centerX+2][centerY+2] {
					count++
				}
			}
		}
	}

	if count > 1 {
		return count
	}

	for i := 1; i <= 2; i++ {
		//Check left diagonals
		if dna[i][adjX1] == dna[i+adjX2][adjX2] {
			if dna[i+1][adjX1+1] == dna[i+2][adjX2-1] {
				count++
			}
		} else if i == 1 {
			//Its necesary to check both of possibilities in the nearest diagonal from  the main diagonal. When i = 1
			if dna[centerY][centerX] == dna[centerY+1][centerX+1] {
				if dna[centerY-1][centerX-1] == dna[centerY+2][centerX+2] {
					count++
				}
			}
		}
	}
	return count
}

func checkDiagonalNegative(dna [][]string) int {
	log.Println("[dna]: ", dna)
	count := 0

	//1 step: Check main diagonal
	count = checkMainDiagonal(dna, count)

	//2 step, check adjacent diagonals
	count = checkAdjacentDiagonal(dna, count)

	log.Println("[count]: ", count)
	return count
}

func checkDiagonalPositive(dna [][]string) int {
	dnaMatrix := utils.ReverseMatrix(dna)
	count := checkDiagonalNegative(dnaMatrix)
	return count
}
