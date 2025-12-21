package p0955

import "testing"

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{strs: []string{"ca", "bb", "ac"}}, 1},
		{"Example 2", args{strs: []string{"xc", "yb", "za"}}, 0},
		{"Example 3", args{strs: []string{"zyx", "wvu", "tsr"}}, 3},
		{"TestCase 89", args{strs: []string{"xga", "xfb", "yfa"}}, 1},
		{"TestCase 125", args{strs: []string{"vdy", "vei", "zvc", "zld"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.args.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
