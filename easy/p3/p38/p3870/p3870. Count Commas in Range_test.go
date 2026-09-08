package p3870

import "testing"

func Test_countCommas(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 1002}, 3},
		{"Example 2", args{n: 998}, 0},
		{"My 1", args{n: 20444}, 19445},
		{"My 2", args{n: 100000}, 99001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCommas(tt.args.n); got != tt.want {
				t.Errorf("countCommas() = %v, want %v", got, tt.want)
			}
		})
	}
}
