package p0786

import (
	"reflect"
	"testing"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{1, 2, 3, 5}, 3}, []int{2, 5}},
		{"Example 2", args{[]int{1, 7}, 1}, []int{1, 7}},
		{"Example 58", args{[]int{1, 811, 1997, 2137, 2803, 3733, 4391, 4441, 4969, 6353, 6991, 7057, 9257, 9323, 11351, 11551, 11821, 11887, 12347, 12743, 13487, 14249, 14869, 15287, 15937, 16411, 18077, 18169, 18397, 18443, 18701, 18839, 19031, 21577, 22669, 23357, 23509, 24329, 24473, 25577, 25931, 25943, 26267, 26833, 27673, 27823, 28051, 28723, 29009, 29207}, 1194}, []int{11551, 11821}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestPrimeFraction(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
