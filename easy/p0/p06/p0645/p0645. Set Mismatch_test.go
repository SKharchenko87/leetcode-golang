package p0645

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{[]int{1, 2, 2, 4}}, []int{2, 3}},
		{"Case 2", args{[]int{1, 1}}, []int{1, 2}},
		{"Case 18", args{[]int{3, 2, 3, 4, 6, 5}}, []int{3, 1}},
		{"Case 29", args{[]int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9}}, []int{2, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNums(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
