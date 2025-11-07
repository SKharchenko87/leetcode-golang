package p1488

import (
	"reflect"
	"testing"
)

func Test_avoidFlood(t *testing.T) {
	type args struct {
		rains []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{rains: []int{1, 2, 3, 4}}, []int{-1, -1, -1, -1}},
		{"Example 2", args{rains: []int{1, 2, 0, 0, 2, 1}}, []int{-1, -1, 2, 1, -1, -1}},
		{"Example 3", args{rains: []int{1, 2, 0, 1, 2}}, []int{}},
		{"TestCase 47", args{rains: []int{0, 1, 1}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avoidFlood(tt.args.rains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avoidFlood() = %v, want %v", got, tt.want)
			}
		})
	}
}
