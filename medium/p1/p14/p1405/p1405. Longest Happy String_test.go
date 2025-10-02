package p1405

import "testing"

func Test_longestDiverseString(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{a: 1, b: 1, c: 7}, "ccaccbcc"},
		{"Example 2", args{a: 7, b: 1, c: 0}, "aabaa"},
		{"TestCase 3", args{a: 1, b: 0, c: 3}, "ccac"},
		{"TestCase 7", args{a: 2, b: 2, c: 1}, "ababc"},
		{"TestCase 24", args{a: 4, b: 4, c: 3}, "aabbccaabbc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDiverseString(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("longestDiverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
