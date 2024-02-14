package p2108

import "testing"

func Test_firstPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{[]string{"abc", "car", "ada", "racecar", "cool"}}, "ada"},
		{"Example 2", args{[]string{"notapalindrome", "racecar"}}, "racecar"},
		{"Example 3", args{[]string{"def", "ghi"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstPalindrome(tt.args.words); got != tt.want {
				t.Errorf("firstPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
