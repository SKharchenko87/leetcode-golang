package p2

import "testing"

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[][]int{{1, 1}, {2, 2}, {3, 3}}}, 0},
		{"Case 1", args{[][]int{{6, 2}, {4, 4}, {2, 6}}}, 2},
		{"Case 1", args{[][]int{{3, 1}, {1, 3}, {1, 1}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPairs(tt.args.points); got != tt.want {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
