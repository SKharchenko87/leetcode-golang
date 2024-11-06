package p0796

import "testing"

func Test_rotateString(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s: "abcde", goal: "cdeab"}, true},
		{"Example 2", args{s: "abcde", goal: "abced"}, false},
		{"TestCase 49", args{s: "defdefdefabcabc", goal: "defdefabcabcdef"}, true},
		{"TestCase 52", args{s: "bbbacddceeb", goal: "ceebbbbacdd"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKVM(t *testing.T) {
	type args struct {
		src string
		trg string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"abcde", "ab"}, 0},
		{"Example 2", args{"abcde", "cd"}, 2},
		{"Example 3", args{"abcde", "d"}, 3},
		{"Example 4", args{"bbbacddceeb", "ceebbbbacddceebbbbacdd"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKVM(tt.args.src, tt.args.trg); got != tt.want {
				t.Errorf("findKVM() = %v, want %v", got, tt.want)
			}
		})
	}
}
