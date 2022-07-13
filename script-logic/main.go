package main

import (
	"fmt"
	"log"
)

func main() {
	dnaRequest := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	dnaMatrix := sliceToMatrix(dnaRequest)

	dnaResponse := checkIfMutant(dnaMatrix)
	fmt.Println("result", dnaResponse)
}

//transformIntoMatrix takes a slice of strings and returns a slice of slices of strings
//Inputs: slice of strings
//Outputs: slice of slices of strings
func sliceToMatrix(slice []string) [][]string {
	var dna [][]string
	for i := 0; i < len(slice); i++ {
		dna = append(dna, convertStringIntoChars(slice[i]))
	}
	return dna
}

//transformIntoMatrix Transpose matrix
//Inputs: slice of strings - Original matrix
//Outputs: slice of strings - Transposed matrix
func transposeMatrix(slice [][]string) [][]string {
	var newSlice [][]string
	for i := 0; i < len(slice[0]); i++ {
		var newRow []string
		for j := 0; j < len(slice); j++ {
			newRow = append(newRow, slice[j][i])
		}
		newSlice = append(newSlice, newRow)
	}
	return newSlice
}

//checkIfMutant is used to check if there is an anomaly pattern in the dna sequence
func checkIfMutant(dna [][]string) bool {
	horizontal := checkHorizontal(dna)
	vertical := checkVertical(dna)
	diagonal := checkDiagonal(dna)
	return diagonal+horizontal+vertical > 1
	//return horizontal+diagonal > 1
}

func checkDiagonal(dna [][]string) int {
	count := 0
	return count
}

func checkHorizontal(dna [][]string) int {
	fmt.Println("[Matrix]", dna)
	count := 0

	for i := 0; i < len(dna); i++ {
		//search hints for possible patterns
		if 3 < len(dna[i]) {
			if dna[i][0] == dna[i][2] {
				if dna[i][0] == dna[i][1] && dna[i][1] == dna[i][3] {
					count++
				}
			}
			if dna[i][3] == dna[i][4] {
				if dna[i][2] == dna[i][3] && dna[i][1] == dna[i][3] {
					count++
				}
			}
		}

	}
	log.Println("[count]", count)
	return count
}

func checkVertical(dna [][]string) int {
	dnaMatrix := transposeMatrix(dna)
	count := checkHorizontal(dnaMatrix)
	return count
}

func checkIf4ConsecutiveChar(dna [][]string) bool {
	count := 0
	for i := 0; i < len(dna); i++ {
		for j := 0; j < len(dna[i]); j++ {
			if dna[i][j] == dna[i][j+1] {
				count++
			}
		}
		if count == 3 {
			return true
		}
		count = 0
	}
	return false
}

func stringToGrid(dna string, col int) []string {
	var newDna []string
	for i := 0; i < len(dna); i += col {
		newDna = append(newDna, dna[i:i+col])
	}
	return newDna
}

func printStringsIntoChars(dna string) {
	for i := 0; i < len(dna); i++ {
		fmt.Println(dna[i])
	}
}

func convertStringIntoChars(dna string) []string {
	var newDna []string
	for i := 0; i < len(dna); i++ {
		newDna = append(newDna, string(dna[i]))
	}
	return newDna
}
