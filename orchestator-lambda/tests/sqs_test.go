package tests

import ("testing"
"github.com/davidpolme/mutant-detector/orchestator-lambda/models"")

func Test_SendToSQS(t *testing.T) {
	type args struct {
		request models.Request
		b       bool
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"", nil, nil},
}
