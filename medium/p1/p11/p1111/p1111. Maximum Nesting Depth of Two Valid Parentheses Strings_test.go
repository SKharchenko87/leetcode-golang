package p1111

import (
	"reflect"
	"testing"
)

func Test_maxDepthAfterSplit(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{seq: "(()())"}, []int{0, 1, 1, 1, 1, 0}},
		{"Example 2", args{seq: "()(())()"}, []int{0, 0, 0, 1, 1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthAfterSplit(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDepthAfterSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
