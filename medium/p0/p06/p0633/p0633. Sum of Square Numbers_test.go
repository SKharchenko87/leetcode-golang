package p0633

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{5}, true},
		{"Example 2", args{3}, false},
		{"Tast 63", args{7}, false},
		{"Tast 62", args{3}, false},
		{"Tast 125", args{8}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
