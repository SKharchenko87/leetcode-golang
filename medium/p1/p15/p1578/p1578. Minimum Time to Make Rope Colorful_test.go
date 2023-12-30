package p1578

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		colors     string
		neededTime []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{"abaac", []int{1, 2, 3, 4, 5}}, 3},
		{"Case 2", args{"abc", []int{1, 2, 3}}, 0},
		{"Case 3", args{"aabaa", []int{1, 2, 3, 4, 1}}, 2},
		{"Case 4", args{"aabaaaa", []int{1, 2, 3, 4, 1, 1, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.colors, tt.args.neededTime); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
