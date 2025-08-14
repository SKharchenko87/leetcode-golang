package p4

import "testing"

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

func Test_maxPathLength(t *testing.T) {
	type args struct {
		coordinates [][]int
		k           int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{coordinates: [][]int{{3, 1}, {2, 2}, {4, 1}, {0, 0}, {5, 3}}, k: 1}, 3},
		{"Example 2", args{coordinates: [][]int{{2, 1}, {7, 0}, {5, 6}}, k: 2}, 2},
		{"TestCase 186", args{coordinates: [][]int{{0, 1}, {0, 3}}, k: 0}, 1},
		{"TestCase 340", args{coordinates: [][]int{{0, 3}, {8, 5}, {6, 8}}, k: 0}, 2},
		//{"TestCase 523", args{coordinates: [][]int{{9, 8}, {2, 1}, {1, 0}, {6, 0}}, k: 0}, 3},
		{"TestCase 475", args{coordinates: [][]int{{5, 0}, {9, 3}, {9, 8}}, k: 0}, 2},
		{"TestCase 573", args{coordinates: [][]int{{9, 8}, {2, 1}, {1, 0}, {6, 0}}, k: 0}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathLength(tt.args.coordinates, tt.args.k); got != tt.want {
				t.Errorf("maxPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Input
//coordinates =
//[[0,1},{0,3]]
//k =
//0
//Use Testcase
//Output
//2
//Expected
//1
