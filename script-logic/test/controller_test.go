package tests

import (
	"testing"

	"github.com/davidpolme/mutant-detector/script-logic/controllers"
	"github.com/davidpolme/mutant-detector/script-logic/utils"
)

func TestCheckMutant(t *testing.T) {
	dnaRequest := []string{"TCACTG", "CCACTA", "AACACA", "CACACA", "CAGTGC", "AGGGTT"}
	//Convert slices to matrix
	dnaMatrix := utils.SliceToMatrix(dnaRequest)
	//Check if there is a pattern

	dnaResponse := controllers.CheckIfMutant(dnaMatrix)
	if dnaResponse {
		t.Errorf("Expected false, got %v", dnaResponse)
	}
}
