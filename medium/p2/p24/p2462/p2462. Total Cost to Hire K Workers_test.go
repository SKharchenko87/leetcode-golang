package p2462

import "testing"

func Test_totalCost(t *testing.T) {
	type args struct {
		costs      []int
		k          int
		candidates int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Case 1", args{[]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4}, 11},
		{"Case 2", args{[]int{1, 2, 4, 1}, 3, 3}, 4},
		{"Case 10", args{[]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2}, 423},
		{"Case 12", args{[]int{28, 35, 21, 13, 21, 72, 35, 52, 74, 92, 25, 65, 77, 1, 73, 32, 43, 68, 8, 100, 84, 80, 14, 88, 42, 53, 98, 69, 64, 40, 60, 23, 99, 83, 5, 21, 76, 34}, 32, 12}, 1407},
		{"Case 20", args{[]int{57, 33, 26, 76, 14, 67, 24, 90, 72, 37, 30}, 11, 2}, 526},
		{"Case 131", args{[]int{60, 87, 94, 42, 12, 60}, 5, 4}, 261},
		{"Case 160", args{[]int{10, 1, 11, 10}, 2, 1}, 11},
		{"Case 161", args{[]int{2, 2, 2, 2, 2, 2, 1, 4, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 7, 2}, 13},
		//ToDo
		//{"My 1", args{[]int{2, 2, 2, 2, 2, 2, 2, 4, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2}, 7, 2}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalCost(tt.args.costs, tt.args.k, tt.args.candidates); got != tt.want {
				t.Errorf("totalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
