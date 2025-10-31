package p1526

import "testing"

func Test_minNumberOperations(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{target: []int{1, 2, 3, 2, 1}}, 3},
		{"Example 2", args{target: []int{3, 1, 1, 2}}, 4},
		{"Example 3", args{target: []int{3, 1, 5, 4, 2}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOperations(tt.args.target); got != tt.want {
				t.Errorf("minNumberOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
