package p1353

import "testing"

func Test_maxEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{events: [][]int{{1, 2}, {2, 3}, {3, 4}}}, 3},
		{"Example 2", args{events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}}, 4},
		{"TestCase 33", args{events: [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}}, 4},
		{"TestCase 34", args{events: [][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.args.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
