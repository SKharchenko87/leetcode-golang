package p3

import (
	"reflect"
	"testing"
)

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resultGrid(t *testing.T) {
	type args struct {
		image     [][]int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Case 1", args{[][]int{{5, 6, 7, 10}, {8, 9, 10, 10}, {11, 12, 13, 10}}, 3}, [][]int{{9, 9, 9, 9}, {9, 9, 9, 9}, {9, 9, 9, 9}}},
		{"Case 2", args{[][]int{{10, 20, 30}, {15, 25, 35}, {20, 30, 40}, {25, 35, 45}}, 12}, [][]int{{25, 25, 25}, {27, 27, 27}, {27, 27, 27}, {30, 30, 30}}},
		{"Case 3", args{[][]int{{5, 6, 7}, {8, 9, 10}, {11, 12, 13}}, 1}, [][]int{{5, 6, 7}, {8, 9, 10}, {11, 12, 13}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultGrid(tt.args.image, tt.args.threshold); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
