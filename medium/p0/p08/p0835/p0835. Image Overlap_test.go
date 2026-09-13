package p0835

import (
	"reflect"
	"testing"
)

func Test_getLines(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"Test0", args{[][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}}, []uint{6, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLines(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestOverlap(t *testing.T) {
	type args struct {
		img1 [][]int
		img2 [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{img1: [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, img2: [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}}, 3},
		{"Example 2", args{img1: [][]int{{1}}, img2: [][]int{{1}}}, 1},
		{"Example 3", args{img1: [][]int{{0}}, img2: [][]int{{0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestOverlap(tt.args.img1, tt.args.img2); got != tt.want {
				t.Errorf("largestOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
