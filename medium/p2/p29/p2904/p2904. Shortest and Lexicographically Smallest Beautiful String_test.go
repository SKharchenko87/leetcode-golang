package p2904

import "testing"

func Test_shortestBeautifulSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "100011001", k: 3}, "11001"},
		{"Example 2", args{s: "1011", k: 2}, "11"},
		{"Example 3", args{s: "000", k: 1}, ""},
		{"TestCase 451", args{s: "111111110010001010", k: 11}, "11111111001000101"},
		{"TestCase 522", args{s: "1100001110111100100", k: 8}, "11101111001"},
		{"TestCase 671", args{s: "001", k: 1}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestBeautifulSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("shortestBeautifulSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
