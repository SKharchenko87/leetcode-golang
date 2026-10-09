package p1541

import "testing"

func Test_minInsertions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "(()))"}, 1},
		{"Example 2", args{s: "())"}, 0},
		{"Example 3", args{s: "))())("}, 3},
		{"TestCase 26", args{s: ")))))))"}, 5},
		{"TestCase 28", args{s: "()()()()()("}, 7},
		{"TestCase 11", args{s: "(((((("}, 12},
		{"TestCase 33", args{s: "))))()))())"}, 4},
		{"TestCase 78", args{s: "))))))((()))(()(()))"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
