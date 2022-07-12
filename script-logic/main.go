package main

import (
	"fmt"
	"log"
)

func main() {
	dnaRequest := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	fmt.Println(dnaRequest)
	dnaResponse := printItemsOfSlice(dnaRequest)
	fmt.Println(dnaResponse)

	newT:=transpose(dnaResponse)
fmt.Println(newT)
}

func convertStringIntoChars(dna string) []string {
	var newDna []string
	for i := 0; i < len(dna); i++ {
		newDna = append(newDna, string(dna[i]))
	}
	return newDna
}

func printItemsOfSlice(slice []string) [][]string {
	//declarevariable 2d array of strings
	var dna [][]string
	for i := 0; i < len(slice); i++ {
		fmt.Println()
		dna = append(dna, convertStringIntoChars(slice[i]))
	}
	return dna
}

func checkIfThereAre4CharsInARow(dna string) bool {
	coincidences := 0
	for i := 0; i < len(dna)-3; i++ {
		if dna[i] == dna[i+1] && dna[i+1] == dna[i+2] && dna[i+2] == dna[i+3] {
			coincidences++
		}
		log.Println("Coincidence:", coincidences, dna[i], dna[i+1], dna[i+2], dna[i+3])
	}
	return coincidences > 1
}

func transpose(slice [][]string) [][]string {
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

func checkIf4ConsecutiveChar(dna []string) bool {
	coincidences := 0

	for i := 0; i < len(dna[0]); i++ {
		if dna[0][i] == dna[1][i] && dna[1][i] == dna[2][i] && dna[2][i] == dna[3][i] {
			coincidences++
		}
		if dna[0][i] == dna[0][i+1] && dna[0][i+1] == dna[0][i+2] && dna[0][i+2] == dna[0][i+3] {
			coincidences++
		}
		log.Println("Coincidence:", coincidences, dna[0][i], dna[1][i], dna[2][i], dna[3][i])
	}
	return coincidences > 1
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
