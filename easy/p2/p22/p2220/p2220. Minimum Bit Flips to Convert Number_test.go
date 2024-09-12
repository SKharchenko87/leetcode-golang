package p2220

import "testing"

func Test_minBitFlips(t *testing.T) {
	type args struct {
		start int
		goal  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{start: 10, goal: 7}, 3},
		{"Example 2", args{start: 3, goal: 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBitFlips(tt.args.start, tt.args.goal); got != tt.want {
				t.Errorf("minBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
