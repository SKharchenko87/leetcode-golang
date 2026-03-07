package p1888

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "111000"}, 2},
		{"Example 2", args{s: "010"}, 0},
		{"Example 3", args{s: "1110"}, 1},
		{"My 1", args{s: "1111000"}, 2},
		{"TestCase", args{s: "01001001101"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.s); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
