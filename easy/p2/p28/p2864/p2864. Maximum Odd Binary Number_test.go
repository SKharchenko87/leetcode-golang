package p2864

import "testing"

func Test_maximumOddBinaryNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"010"}, "001"},
		{"Example 2", args{"0101"}, "1001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumOddBinaryNumber(tt.args.s); got != tt.want {
				t.Errorf("maximumOddBinaryNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
