package p0140

import (
	"reflect"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}}, []string{"cats and dog", "cat sand dog"}},
		{"Example 2", args{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}}, []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
		{"Example 3", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, []string{}},
		{"Example 3", args{"a", []string{"a"}}, []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPositions(t *testing.T) {
	type args struct {
		s         string
		substring string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"E1", args{"saskklsasas", "sas"}, []int{0, 6, 8}},
		{"E1", args{"saskklsasas", "x"}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findPositions(tt.args.s, tt.args.substring); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findPositions() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
