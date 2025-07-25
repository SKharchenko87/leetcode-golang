package p2942

import (
	"reflect"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	type args struct {
		words []string
		x     byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{words: []string{"leet", "code"}, x: 'e'}, []int{0, 1}},
		{"Example 2", args{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: 'a'}, []int{0, 2}},
		{"Example 3", args{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: 'z'}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordsContaining(tt.args.words, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
