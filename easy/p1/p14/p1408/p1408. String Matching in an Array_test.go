package p1408

import (
	"reflect"
	"testing"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{words: []string{"mass", "as", "hero", "superhero"}}, []string{"as", "hero"}},
		{"Example 2", args{words: []string{"leetcode", "et", "code"}}, []string{"et", "code"}},
		{"Example 3", args{words: []string{"blue", "green", "bu"}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringMatching(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
