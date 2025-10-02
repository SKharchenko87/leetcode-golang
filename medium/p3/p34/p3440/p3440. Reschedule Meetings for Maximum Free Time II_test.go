package p3440

import "testing"

func Test_maxFreeTime(t *testing.T) {
	type args struct {
		eventTime int
		startTime []int
		endTime   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{eventTime: 5, startTime: []int{1, 3}, endTime: []int{2, 5}}, 2},
		{"Example 2", args{eventTime: 10, startTime: []int{0, 7, 9}, endTime: []int{1, 8, 10}}, 7},
		{"Example 3", args{eventTime: 10, startTime: []int{0, 3, 7, 9}, endTime: []int{1, 4, 8, 10}}, 6},
		{"Example 4", args{eventTime: 5, startTime: []int{0, 1, 2, 3, 4}, endTime: []int{1, 2, 3, 4, 5}}, 0},
		{"TestCase 100", args{eventTime: 41, startTime: []int{17, 24}, endTime: []int{19, 25}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreeTime(tt.args.eventTime, tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
