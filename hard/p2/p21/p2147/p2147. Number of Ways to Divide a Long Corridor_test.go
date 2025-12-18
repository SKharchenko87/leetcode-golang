package p2147

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		corridor string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{corridor: "SSPPSPS"}, 3},
		{"Example 2", args{corridor: "PPSPSP"}, 1},
		{"Example 3", args{corridor: "S"}, 0},
		{"My 1", args{corridor: "SSPPSPSS"}, 0},
		{"TestCase 247", args{corridor: "SS"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.corridor); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
