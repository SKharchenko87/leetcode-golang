package p3208

import "testing"

func Test_numberOfAlternatingGroups(t *testing.T) {
	type args struct {
		colors []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{colors: []int{0, 1, 0, 1, 0}, k: 3}, 3},
		{"Example 2", args{colors: []int{0, 1, 0, 0, 1, 0, 1}, k: 6}, 2},
		{"Example 3", args{colors: []int{1, 1, 0, 1}, k: 4}, 0},
		{"TestCase 351", args{colors: []int{0, 0, 0, 1}, k: 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfAlternatingGroups(tt.args.colors, tt.args.k); got != tt.want {
				t.Errorf("numberOfAlternatingGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
