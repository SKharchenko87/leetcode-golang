package p2901

import (
	"reflect"
	"testing"
)

func Test_getWordsInLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{words: []string{"bab", "dab", "cab"}, groups: []int{1, 2, 2}}, []string{"bab", "cab"}},
		{"Example 2", args{words: []string{"a", "b", "c", "d"}, groups: []int{1, 2, 3, 4}}, []string{"a", "b", "c", "d"}},
		{"TestCase 302", args{words: []string{"ca", "cb", "bcd", "bb", "ddc"}, groups: []int{2, 4, 2, 5, 1}}, []string{"ca", "cb", "bb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWordsInLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWordsInLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
