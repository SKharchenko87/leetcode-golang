package p2337

import "testing"

func Test_canChange(t *testing.T) {
	type args struct {
		start  string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{start: "_L__R__R_", target: "L______RR"}, true},
		{"Example 2", args{start: "R_L_", target: "__LR"}, false},
		{"Example 3", args{start: "_R", target: "R_"}, false},
		{"TestCase 6", args{start: "_LL__R__R_", target: "L___L___RR"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canChange(tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("canChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
