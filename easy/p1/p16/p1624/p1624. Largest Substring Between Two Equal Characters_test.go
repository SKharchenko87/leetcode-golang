package p1624

import "testing"

func Test_maxLengthBetweenEqualCharacters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{"aa"}, 0},
		{"Case 1", args{"abca"}, 2},
		{"Case 1", args{"cbzxy"}, -1},
		{"Case 1", args{"mgntdygtxrvxjnwksqhxuxtrv"}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLengthBetweenEqualCharacters(tt.args.s); got != tt.want {
				t.Errorf("maxLengthBetweenEqualCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
