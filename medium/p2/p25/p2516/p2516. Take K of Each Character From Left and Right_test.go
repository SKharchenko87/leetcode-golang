package p2516

import "testing"

func Test_takeCharacters(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "aabaaaacaabc", k: 2}, 8},
		{"Example 2", args{s: "a", k: 1}, -1},
		{"Example 6", args{s: "abc", k: 1}, 3},
		{"Example 131", args{s: "acba", k: 1}, 3},
		{"Example 131", args{s: "aabbccca", k: 2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := takeCharacters(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("takeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
