package p2078

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		colors []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{colors: []int{1, 1, 1, 6, 1, 1, 1}}, 3},
		{"Example 2", args{colors: []int{1, 8, 3, 8, 3}}, 4},
		{"Example 3", args{colors: []int{0, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.colors); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
