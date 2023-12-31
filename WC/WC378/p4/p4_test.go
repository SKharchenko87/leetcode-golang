package p4

import (
	"reflect"
	"testing"
)

func Test_canMakePalindromeQueries(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{"Case 1", args{"abcabc", [][]int{{1, 1, 3, 5}, {0, 2, 5, 5}}}, []bool{true, true}},
		{"Case 2", args{"abbcdecbba", [][]int{{0, 2, 7, 9}}}, []bool{false}},
		{"Case 3", args{"acbcab", [][]int{{1, 2, 4, 5}}}, []bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakePalindromeQueries(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canMakePalindromeQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
