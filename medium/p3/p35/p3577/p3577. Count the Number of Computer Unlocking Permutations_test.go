package p3577

import "testing"

func Test_countPermutations(t *testing.T) {
	type args struct {
		complexity []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{complexity: []int{1, 2, 3}}, 2},
		{"Example 2", args{complexity: []int{3, 3, 3, 4, 4, 4}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPermutations(tt.args.complexity); got != tt.want {
				t.Errorf("countPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
