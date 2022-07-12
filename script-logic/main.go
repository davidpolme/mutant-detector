package main

import (
	"fmt"
	"strings"
)

func main() {
	dnaRequest := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	dna := strings.Join(dnaRequest, "")
	newDna:=stringToGrid(dna, 6)
	fmt.Println(newDna)
}

func stringToGrid(dna string, col int)[]string{
	var newDna []string
	for i:=0; i<len(dna); i+=col{
		newDna = append(newDna, dna[i:i+col])
	}
	return newDna
}
