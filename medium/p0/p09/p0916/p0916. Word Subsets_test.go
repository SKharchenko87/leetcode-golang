package p0916

import (
	"reflect"
	"testing"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"e", "o"}}, []string{"facebook", "google", "leetcode"}},
		{"Example 2", args{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"l", "e"}}, []string{"apple", "google", "leetcode"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordSubsets(tt.args.words1, tt.args.words2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_wordSubsets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"})
	}
}

func Benchmark_wordSubsets0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordSubsets0([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"})
	}
}
