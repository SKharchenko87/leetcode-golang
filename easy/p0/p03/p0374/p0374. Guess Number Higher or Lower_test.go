package p0374

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		x int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{10, 6}, 6},
		{"Case 1", args{3, 2}, 2},
		{"Case 1", args{1000, 50}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
