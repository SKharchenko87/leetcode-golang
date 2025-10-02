package p0179

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{nums: []int{10, 2}}, "210"},
		{"Example 2", args{nums: []int{3, 30, 34, 5, 9}}, "9534330"},
		{"Example 2", args{nums: []int{3, 300, 44, 5, 9}}, "95443300"},
		{"Example 2", args{nums: []int{123, 34, 5, 9}}, "9534123"},
		{"Example 2", args{nums: []int{139, 34, 5, 9}}, "9534139"},
		{"Example 2", args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}, "9876543210"},
		{"Example 2", args{nums: []int{0, 0}}, "0"},
		{"Example 2", args{nums: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}}, "9876543210"},
		{"Example 2", args{nums: []int{111311, 1113}}, "1113111311"},
		{"Example 2", args{nums: []int{999999991, 9}}, "9999999991"},
		{"Example 2", args{nums: []int{1000000000, 1000000000}}, "10000000001000000000"},
		{"Example 2", args{nums: []int{41, 23, 87, 55, 50, 53, 18, 9, 39, 63, 35, 33, 54, 25, 26, 49, 74, 61, 32, 81, 97, 99, 38, 96, 22, 95, 35, 57, 80, 80, 16, 22, 17, 13, 89, 11, 75, 98, 57, 81, 69, 8, 10, 85, 13, 49, 66, 94, 80, 25, 13, 85, 55, 12, 87, 50, 28, 96, 80, 43, 10, 24, 88, 52, 16, 92, 61, 28, 26, 78, 28, 28, 16, 1, 56, 31, 47, 85, 27, 30, 85, 2, 30, 51, 84, 50, 3, 14, 97, 9, 91, 90, 63, 90, 92, 89, 76, 76, 67, 55}}, "99999897979696959492929190908989888878785858585848181808080807876767574696766636361615757565555555453525150505049494743413938353533332313030282828282726262525242322222181716161614131313121111010"},
		{"Example 2", args{nums: []int{11, 10, 10, 1, 2}}, "21111010"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numToSlice(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test", args{num: 10}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numToSlice(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_XXX(t *testing.T) {
	fmt.Printf("%b\n", 1)
	fmt.Printf("%b\n", -1)
}
