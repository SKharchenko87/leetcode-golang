package p1641

import "testing"

func Test_countVowelStrings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 1}, 5},
		{"Example 2", args{n: 2}, 15},
		{"Example 3", args{n: 33}, 66045},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelStrings(tt.args.n); got != tt.want {
				t.Errorf("countVowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
