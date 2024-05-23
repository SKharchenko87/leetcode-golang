package p0131

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]string
	}{
		{"Example 1", args{"aab"}, [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"Example 2", args{"a"}, [][]string{{"a"}}},
		//{"Example 2*", args{"aaa"}, [][]string{{"aaa"}}},
		{"Example 1*", args{"abb"}, [][]string{{"a", "b", "b"}, {"a", "bb"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := partition(tt.args.s); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("partition() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
