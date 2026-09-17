package p2472

import "testing"

func Test_maxPalindromes(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "abaccdbbd", k: 3}, 2},
		{"Example 2", args{s: "adbcda", k: 2}, 0},
		{"TestCase 42", args{s: "iqqibcecvrbxxj", k: 1}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPalindromes(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxPalindromes() = %v, want %v", got, tt.want)
			}
		})
	}
}
