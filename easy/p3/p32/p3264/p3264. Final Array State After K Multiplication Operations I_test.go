package p3264

import (
	"reflect"
	"testing"
)

func Test_getFinalState(t *testing.T) {
	type args struct {
		nums       []int
		k          int
		multiplier int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{2, 1, 3, 5, 6}, k: 5, multiplier: 2}, []int{8, 4, 6, 5, 6}},
		{"Example 2", args{nums: []int{1, 2}, k: 3, multiplier: 4}, []int{16, 8}},
		{"TestCase 251", args{nums: []int{2, 1}, k: 1, multiplier: 2}, []int{2, 2}},
		{"TestCase 733", args{nums: []int{38, 6}, k: 6, multiplier: 2}, []int{152, 96}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFinalState(tt.args.nums, tt.args.k, tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFinalState() = %v, want %v", got, tt.want)
			}
		})
	}
}
