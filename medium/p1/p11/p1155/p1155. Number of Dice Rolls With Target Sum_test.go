package p1155

import "testing"

func Test_numRollsToTarget(t *testing.T) {
	type args struct {
		n      int
		k      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{1, 6, 3}, 1},
		{"Case 2", args{2, 6, 7}, 6},
		{"Case 2", args{3, 6, 11}, 27},
		{"Case 3", args{30, 30, 500}, 222616187},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRollsToTarget(tt.args.n, tt.args.k, tt.args.target); got != tt.want {
				t.Errorf("numRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
