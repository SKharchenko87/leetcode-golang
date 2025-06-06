package p0763

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{s: "ababcbacadefegdehijhklij"}, []int{9, 7, 8}},
		{"Example 2", args{s: "eccbbbbdec"}, []int{10}},
		{"TestCase 0", args{s: "ab"}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
