package tests

import (
	"fmt"
	"testing"

	"github.com/davidpolme/mutant-detector/lambda-worker/controller"
)

func TestCheckMutant(t *testing.T) {
	var tests = []struct {
		dnaRequest string
		want       bool
	}{
		{"AGGGGTCAGTGCCACACAAACACACAACTATCACTG", true},
		{"AGAGGTCAGTTCCAATCAAATACACCACTATCACTG", true},
		{"AGAGGACTGTTCCATTCAAATTCTCTACTAACACTA", true},// TODO: middle diagonal
		{"TGAGGACTGTACCAATCAAATACTCTACAATCACTA", true},
		{"CAGTTCCAATCAAATACACTACAAACACTGAGAGGT", true},
		{"ACACTGTGAGGACAGTTCCAATCAAATACACTACAA", true},
		{"TGAGGACTGTACAAATCTAATGTTCTATAATCTATA", true},
		{"AGGGTTCAGTGCCACACAAACACACCACTATCACTG", false},
		{"AGAGGTCAGTGCCAATCAAATACACCACAGTCACTA", false},
		{"AGAGGTAGGTTCCAGTCATTAACACTAATATCACAG", false},
		{"AGAGGTAGGTACCAATCATTGCCACTCATATCACAG", false},
		{"AGAGGTAGGTACCAATGATTTCCACTCATATCACAG", false},
		{"AGAGGTAGGTTCCGAGCATTAAGACTAATGTCACAG", false},
		{"AGAGGTAGTGTCCCGGCATGAAGACTAATGTCACAG", false},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%s,%t", tt.dnaRequest, tt.want)
		t.Run(testname, func(t *testing.T) {
			dnaResponse := controller.CheckMutant(tt.dnaRequest)
			if dnaResponse != tt.want {
				t.Errorf("got %t, want %t", dnaResponse, tt.want)
			}
		})
	}
}
