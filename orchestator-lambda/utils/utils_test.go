package utils

import (
	"testing"

	"github.com/davidpolme/mutant-detector/orchestator-lambda/models"
)

func Test_ValidateRequest(t *testing.T) {
	type args struct {
		request models.Request
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{ //Test 1
			name: "Test 1",
			args: args{request: models.Request{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}},
			want: nil,
		},
		{ //Test 2
			name: "Test 2",
			args: args{request: models.Request{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}},
			want: nil,
		},
		{ //Test 3
			name: "Test 3",
			args: args{request: models.Request{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateRequest(tt.args.request); got != tt.want {
				t.Errorf("ValidateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_IsMutant(t *testing.T){
	type args struct {
		request models.Request
	}
	tests := []struct {
		name string
		args args
		want models.Response
		want1 error
	}{
		{ //Test 1
			name: "Test 1",
			args: args{request: models.Request{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}},
			want: models.Response{IsMutant: true},
			want1: nil,
		},
		{ //Test 2
			name: "Test 2",
			args: args{request: models.Request{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}},
			want: models.Response{IsMutant: true},
			want1: nil,
		},
		{ //Test 3
			name: "Test 3",
			args: args{request: models.Request{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}},
			want: models.Response{IsMutant: true},
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsMutant(tt.args.request)
			if got != tt.want {
				t.Errorf("IsMutant() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsMutant() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}



func Test_ContainsValidChars(t *testing.T) {
	type args struct {
		dna []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{ //Test 1
			name: "Test 1",
			args: args{dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			want: true,
		},
		{ //Test 2
			name: "Test 2",
			args: args{dna: []string{"ATSCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			want: false,
		},
		{ //Test 3
			name: "Test 3",
			args: args{dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTM"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsValidChars(tt.args.dna); got != tt.want {
				t.Errorf("ContainsValidChars() = %v, want %v", got, tt.want)
			}
		})
	}
}