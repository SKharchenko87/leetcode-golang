package p3871

import "testing"

func Test_countCommas1(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{n: 1002}, 3},
		{"Example 2", args{n: 998}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCommas(tt.args.n); got != tt.want {
				t.Errorf("countCommas() = %v, want %v", got, tt.want)
			}
		})
	}
}
