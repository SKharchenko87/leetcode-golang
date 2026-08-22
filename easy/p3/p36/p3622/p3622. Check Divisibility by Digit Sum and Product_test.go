package p3622

import "testing"

func Test_checkDivisibility(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example 1", args: args{n: 99}, want: true},
		{name: "Example 2", args: args{n: 23}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDivisibility(tt.args.n); got != tt.want {
				t.Errorf("checkDivisibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
