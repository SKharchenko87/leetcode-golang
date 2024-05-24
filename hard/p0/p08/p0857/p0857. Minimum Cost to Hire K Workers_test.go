package p0857

import "testing"

func Test_mincostToHireWorkers(t *testing.T) {
	type args struct {
		quality []int
		wage    []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{[]int{10, 20, 5}, []int{70, 50, 30}, 2}, 105.00000},
		{"Example 2", args{[]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3}, 30.66667},
		{"Example 18", args{[]int{100, 65, 91, 12, 37, 4, 58, 100, 19, 5, 69, 89, 92, 31, 16, 25, 45, 63, 40, 16}, []int{278, 119, 197, 374, 408, 198, 56, 246, 308, 459, 217, 334, 189, 460, 106, 451, 120, 129, 104, 498}, 8}, 1477.33333},
		{"Example 33", args{[]int{25, 68, 35, 62, 52, 57, 35, 83, 40, 51}, []int{147, 97, 251, 129, 438, 443, 120, 366, 362, 343}, 6}, 1979.31429},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostToHireWorkers(tt.args.quality, tt.args.wage, tt.args.k); got != tt.want {
				t.Errorf("mincostToHireWorkers() = %v, want %v", got, tt.want)
			}
		})
	}
}
