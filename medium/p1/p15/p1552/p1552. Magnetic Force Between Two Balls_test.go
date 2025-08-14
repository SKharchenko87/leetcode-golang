package p1552

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		position []int
		m        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{position: []int{1, 2, 3, 4, 7}, m: 3}, 3},
		{"Example 2", args{position: []int{5, 4, 3, 2, 1, 1000000000}, m: 2}, 999999999},
		{"Testcase 26", args{position: []int{79, 74, 57, 22}, m: 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.position, tt.args.m); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
