package p1926

import "testing"

func Test_nearestExit(t *testing.T) {
	type args struct {
		maze     [][]byte
		entrance []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}}, 1},
		{"Case 2", args{[][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}}, 2},
		{"Case 3", args{[][]byte{{'.', '+'}}, []int{0, 0}}, -1},
		{"Case 3", args{[][]byte{{'.', '.'}, {'.', '.'}}, []int{0, 0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestExit(tt.args.maze, tt.args.entrance); got != tt.want {
				t.Errorf("nearestExit() = %v, want %v", got, tt.want)
			}
		})
	}
}
