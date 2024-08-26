package p0650

import (
	"reflect"
	"testing"
)

func Test_minSteps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 3}, 3},
		{"Example 2", args{n: 1}, 0},
		{"My 18", args{n: 18}, 8},
		{"My 12", args{n: 12}, 7},
		{"My 20", args{n: 20}, 9},
		{"My 17", args{n: 17}, 17},
		{"My 512", args{n: 512}, 18},
		{"My 1000", args{n: 1000}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 3", args{n: 3}, []int{3}},
		{"Test 1", args{n: 1}, []int{}},
		{"Test 2", args{n: 2}, []int{2}},
		{"Test 18", args{n: 18}, []int{2, 3, 3}},
		{"Test 12", args{n: 12}, []int{2, 2, 3}},
		{"Test 20", args{n: 20}, []int{2, 2, 5}},
		{"Test 17", args{n: 17}, []int{17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberDivisors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
