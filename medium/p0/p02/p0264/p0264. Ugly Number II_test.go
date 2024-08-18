package p0264

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 10}, 12},
		{"Example 2", args{n: 1}, 1},
		{"Example 7", args{n: 7}, 8},
		{"Example 1000", args{n: 1000}, 51200000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Example 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run()
		})
	}
}
