package p0049

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"Example 1", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{"Example 2", args{[]string{""}}, [][]string{{""}}},
		{"Example 3", args{[]string{"a"}}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
