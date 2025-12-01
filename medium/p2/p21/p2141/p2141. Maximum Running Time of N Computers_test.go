package p2141

import "testing"

func Test_maxRunTime(t *testing.T) {
	type args struct {
		n         int
		batteries []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{n: 2, batteries: []int{3, 3, 3}}, 4},
		{"Example 2", args{n: 2, batteries: []int{1, 1, 1, 1}}, 2},
		{"My 1", args{n: 3, batteries: []int{1, 1, 3, 5}}, 2},
		{"My 2", args{n: 2, batteries: []int{1, 1, 3, 5}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRunTime(tt.args.n, tt.args.batteries); got != tt.want {
				t.Errorf("maxRunTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
