package p2379

import "testing"

func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{blocks: "WBBWWBBWBW", k: 7}, 3},
		{"Example 2", args{blocks: "WBWBBBW", k: 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRecolors(tt.args.blocks, tt.args.k); got != tt.want {
				t.Errorf("minimumRecolors() = %v, want %v", got, tt.want)
			}
		})
	}
}
