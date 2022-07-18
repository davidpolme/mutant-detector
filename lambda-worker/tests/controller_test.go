package tests

import (
	"testing"

	"github.com/davidpolme/mutant-detector/lambda-worker/controller"
)

func TestCheckMutant(t *testing.T) {
	dnaRequest := "AGGGTTCAGTGCCACACAAACACACCACTATCACTG"
	dnaResponse := controller.CheckMutant(dnaRequest)
	if !dnaResponse {
		t.Errorf("Expected true, got %v", dnaResponse)
	}
}
