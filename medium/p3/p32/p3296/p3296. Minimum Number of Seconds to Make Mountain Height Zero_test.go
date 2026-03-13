package p3296

import "testing"

func Test_minNumberOfSeconds(t *testing.T) {
	type args struct {
		mountainHeight int
		workerTimes    []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{mountainHeight: 4, workerTimes: []int{2, 1, 1}}, 3},
		{"Example 2", args{mountainHeight: 10, workerTimes: []int{3, 2, 2, 4}}, 12},
		{"Example 3", args{mountainHeight: 5, workerTimes: []int{1}}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfSeconds(tt.args.mountainHeight, tt.args.workerTimes); got != tt.want {
				t.Errorf("minNumberOfSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
