package p1043

import "testing"

func Test_maxSumAfterPartitioning(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{1, 15, 7, 9, 2, 5, 10}, 3}, 84},
		{"Case 2", args{[]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4}, 83},
		{"Case 3", args{[]int{1}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumAfterPartitioning(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("maxSumAfterPartitioning() = %v, want %v", got, tt.want)
			}
		})
	}
}
