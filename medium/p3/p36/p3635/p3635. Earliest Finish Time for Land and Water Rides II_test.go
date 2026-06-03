package p3635

import "testing"

func Test_earliestFinishTime(t *testing.T) {
	type args struct {
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{landStartTime: []int{2, 8}, landDuration: []int{4, 1}, waterStartTime: []int{6}, waterDuration: []int{3}}, 9},
		{"Example 2", args{landStartTime: []int{5}, landDuration: []int{3}, waterStartTime: []int{1}, waterDuration: []int{10}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earliestFinishTime(tt.args.landStartTime, tt.args.landDuration, tt.args.waterStartTime, tt.args.waterDuration); got != tt.want {
				t.Errorf("earliestFinishTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
