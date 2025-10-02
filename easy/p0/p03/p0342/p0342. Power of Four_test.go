package p0342

import "testing"

func Test_isPowerOfFour1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 16}, true},
		{"Example 2", args{n: 5}, false},
		{"Example 3", args{n: 1}, true},
		{"TesCase 1048", args{n: 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
