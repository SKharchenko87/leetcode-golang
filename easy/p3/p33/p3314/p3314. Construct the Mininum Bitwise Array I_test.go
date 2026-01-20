package p3314

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_minBitwiseArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{2, 3, 5, 7}}, []int{-1, 1, 4, 3}},
		{"Example 2", args{nums: []int{11, 13, 31}}, []int{9, 12, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBitwiseArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minBitwiseArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test1(t *testing.T) {
	for i := 0; i < 32; i++ {
		fmt.Printf("%b \t\t %d\n", i, i)
	}
}
