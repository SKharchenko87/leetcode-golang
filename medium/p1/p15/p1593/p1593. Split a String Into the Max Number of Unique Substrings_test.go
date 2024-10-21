package p1593

import "testing"

func Test_maxUniqueSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "ababccc"}, 5},
		{"Example 2", args{s: "aba"}, 2},
		{"Example 3", args{s: "aa"}, 1},
		{"TestCase 91", args{s: "wwwzfvedwfvhsww"}, 11},
		// wwwzfvedwfvhsww
		// www z f v e d w fv h s ww
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUniqueSplit(tt.args.s); got != tt.want {
				t.Errorf("maxUniqueSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
