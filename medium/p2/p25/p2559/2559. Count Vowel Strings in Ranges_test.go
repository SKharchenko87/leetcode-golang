package p2559

import (
	"reflect"
	"testing"
)

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words   []string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{words: []string{"aba", "bcb", "ece", "aa", "e"}, queries: [][]int{{0, 2}, {1, 4}, {1, 1}}}, []int{2, 3, 0}},
		{"Example 2", args{words: []string{"a", "e", "i"}, queries: [][]int{{0, 2}, {0, 1}, {2, 2}}}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelStrings(tt.args.words, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
