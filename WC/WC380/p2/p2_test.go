package p2

import (
	"reflect"
	"testing"
)

func Test_beautifulIndices(t *testing.T) {
	type args struct {
		s string
		a string
		b string
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{"isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15}, []int{16, 33}},
		{"Case 2", args{"abcd", "a", "a", 4}, []int{0}},
		{"Case 2", args{"ocmm", "m", "oc", 3}, []int{2, 3}},
		{"Case 2", args{"aaaa", "a", "aa", 0}, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulIndices(tt.args.s, tt.args.a, tt.args.b, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("beautifulIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
