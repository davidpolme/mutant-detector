package tests

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/davidpolme/mutant-detector/detector_lambda/handlers"
	"github.com/davidpolme/mutant-detector/detector_lambda/models"
)

func Test_handlers_MyHandler(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.Response
		wantErr bool
	}{
		{
			name: "Test_handlers_MyHandler",
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`,
				},
			},
			want: models.Response{
				IsMutant: true,
			},
			wantErr: false,
		},
		{
			name: "Test_handlers_MyHandler",
			args: args{
				request: events.APIGatewayProxyRequest{
					//TCACTG
					Body: `{"dna":["AGGGTT","CAGTGC","CACACA","AACACA","CCACTA","TCACTG"]}`,
				},
			},
			want: models.Response{
				IsMutant: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.MyHandler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("MyHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MyHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
