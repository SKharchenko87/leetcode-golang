package p1531

import "testing"

func Test_getLengthOfOptimalCompression(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{"aaabcccd", 2}, 4},
		{"Case 2", args{"aabbaa", 2}, 2},
		{"Case 3", args{"aaaaaaaaaaa", 0}, 3},
		{"Case 4", args{"aabbaacd", 2}, 4},
		{"Case 5", args{"aabbbacd", 3}, 4},
		{"Case 6", args{"aabbbacd", 4}, 3},
		{"Case 7", args{"aabbbacd", 5}, 2},
		{"Case 8", args{"aabbbacd", 6}, 2},
		{"Case 46", args{"bbabbbabbbbcbb", 4}, 3},
		{"Case 47", args{"bababbaba", 1}, 7},
		{"Case 47", args{"cdacddcaaacbc", 7}, 4},
		{"Case 66", args{"a", 1}, 0},
		{"Case 71", args{"llllllllllttttttttt", 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLengthOfOptimalCompression(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getLengthOfOptimalCompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
