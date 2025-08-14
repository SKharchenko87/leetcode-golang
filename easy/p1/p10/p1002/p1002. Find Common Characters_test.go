package p1002

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{words: []string{"bella", "label", "roller"}}, []string{"e", "l", "l"}},
		{"Example 2", args{words: []string{"cool", "lock", "cook"}}, []string{"c", "o"}},
		{"Example 2", args{words: []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
