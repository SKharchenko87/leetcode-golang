package p2402

import "testing"

func Test_mostBooked(t *testing.T) {
	type args struct {
		n        int
		meetings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 2, meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}}, 0},
		{"Example 2", args{n: 3, meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}}, 1},
		{"TestCase 68", args{n: 4, meetings: [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}}, 0},
		{"TestCase 79", args{n: 4, meetings: [][]int{{48, 49}, {22, 30}, {13, 31}, {31, 46}, {37, 46}, {32, 36}, {25, 36}, {49, 50}, {24, 34}, {6, 41}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostBooked(tt.args.n, tt.args.meetings); got != tt.want {
				t.Errorf("mostBooked() = %v, want %v", got, tt.want)
			}
		})
	}
}
