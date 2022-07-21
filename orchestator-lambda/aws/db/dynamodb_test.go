package db

import (
	"testing"

	"github.com/davidpolme/mutant-detector/orchestator-lambda/models"
)

func Test_SendToDynamoDB(t *testing.T) {
	type args struct {
		request models.Request
		b       bool
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{ //Test 1
			name: "Test 1",
			args: args{request: models.Request{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
				b: true},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SendToDynamoDB(tt.args.request, tt.args.b); got != tt.want {
				t.Errorf("SendToDynamoDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
