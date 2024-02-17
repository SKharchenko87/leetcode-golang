package p1481

import "testing"

func Test_findLeastNumOfUniqueInts(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{5, 5, 4}, 1}, 1},
		{"Example 2", args{[]int{4, 3, 1, 1, 3, 3, 2}, 3}, 2},
		{"Example 2", args{[]int{2, 4, 1, 8, 3, 5, 1, 3}, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLeastNumOfUniqueInts(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findLeastNumOfUniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
