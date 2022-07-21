package utils

import "testing"

func Test_IsMatrix(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{ //Test 1
			name: "Test 1",
			args: args{ss: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			want: true,
		},
		{ //Test 2
			name: "Test 2",
			args: args{ss: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG", "ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			want: false,
		},
		{ //Test 3
			name: "Test 3",
			args: args{ss: []string{"ATGCGA", "CAGTGC", "TTAATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatrix(tt.args.ss); got != tt.want {
				t.Errorf("IsMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
