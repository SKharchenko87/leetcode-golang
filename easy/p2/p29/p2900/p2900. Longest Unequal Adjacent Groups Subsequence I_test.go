package p2900

import (
	"reflect"
	"testing"
)

func Test_getLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{words: []string{"e", "a", "b"}, groups: []int{0, 0, 1}}, []string{"e", "b"}},
		{"Example 2", args{words: []string{"a", "b", "c", "d"}, groups: []int{1, 0, 1, 1}}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
