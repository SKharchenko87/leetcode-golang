package p1701

import "testing"

func Test_averageWaitingTime(t *testing.T) {
	type args struct {
		customers [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{customers: [][]int{{1, 2}, {2, 5}, {4, 3}}}, 5.00000},
		{"Example 2", args{customers: [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}}, 3.25000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageWaitingTime(tt.args.customers); got != tt.want {
				t.Errorf("averageWaitingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
