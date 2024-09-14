package p1402

import "testing"

func Test_maxSatisfaction(t *testing.T) {
	type args struct {
		satisfaction []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{satisfaction: []int{-1, -8, 0, 5, -9}}, 14},
		{"Example 2", args{satisfaction: []int{4, 3, 2}}, 20},
		{"Example 3", args{satisfaction: []int{-1, -4, -5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfaction(tt.args.satisfaction); got != tt.want {
				t.Errorf("maxSatisfaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
