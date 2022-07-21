package handlers

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/davidpolme/mutant-detector/orchestator-lambda/models"
)

func Test_MyHandler(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name string
		args args
		want models.Response
		err  error
	}{
		{ //Test 1
			name: "Test 1",
			args: args{request: events.APIGatewayProxyRequest{Body: `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`}},
			want: models.Response{IsMutant: true},
			err:  nil,
		},
		{ //Test 2
			name: "Test 2",
			args: args{request: events.APIGatewayProxyRequest{Body: `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`}},
			want: models.Response{IsMutant: true},
			err:  nil,
		},
		{ //Test 3
			name: "Test 3",
			args: args{request: events.APIGatewayProxyRequest{Body: `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`}},
			want: models.Response{IsMutant: true},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MyHandler(tt.args.request)
			if err != tt.err {
				t.Errorf("MyHandler() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
