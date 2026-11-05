package p1385

import "testing"

func Test_findTheDistanceValue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr1 []int
		arr2 []int
		d    int
		want int
	}{
		{"Example 1", []int{4, 5, 8}, []int{10, 9, 1, 8}, 2, 2},
		{"Example 2", []int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3, 2},
		{"Example 3", []int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTheDistanceValue(tt.arr1, tt.arr2, tt.d)
			if got != tt.want {
				t.Errorf("findTheDistanceValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
