package p2807

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{head: []int{18, 6, 10, 3}}, []int{18, 6, 6, 2, 10, 1, 3}},
		{"Example 2", args{head: []int{7}}, []int{7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greatestCommonDivisor(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5/3", args{a: 5, b: 3}, 1},
		{"6/3", args{a: 6, b: 3}, 3},
		{"10/4", args{a: 10, b: 4}, 2},
		{"17/16", args{a: 17, b: 16}, 1},
		{"32/3", args{a: 32, b: 3}, 1},
		{"21/14", args{a: 21, b: 14}, 7},
		{"34/24", args{a: 34, b: 24}, 2},
		{"95/50", args{a: 95, b: 50}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greatestCommonDivisor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("greatestCommonDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}
