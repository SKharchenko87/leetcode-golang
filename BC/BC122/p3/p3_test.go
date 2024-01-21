package p3

import "testing"

func Test_minimumArrayLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//нечетное / нечетное = четное
		//четное / четное = четное
		//четное / нечетное = нечетное
		//нечетное / четное = нечетное
		{"Case 1", args{[]int{1, 4, 3, 1}}, 1},
		{"Case 2", args{[]int{5, 5, 5, 10, 5}}, 2},
		{"Case 3", args{[]int{2, 3, 4}}, 1},
		{"Case 3", args{[]int{1, 1, 5, 5, 5, 6}}, 1},
		{"Case 3", args{[]int{1, 1, 1, 5, 5, 5}}, 2},
		{"Case 3", args{[]int{1, 1, 1, 1, 5, 5, 5}}, 2},
		{"Case 3", args{[]int{1, 1, 1, 1, 1, 5, 5, 5}}, 3},
		{"Case 3", args{[]int{1, 1, 1, 1, 1, 1, 5, 5, 5}}, 3},
		{"Case 3", args{[]int{1, 1, 1, 1, 1, 1, 4, 4, 4}}, 3},
		{"Case 3", args{[]int{3, 3, 3, 3, 3, 3, 5, 5, 5}}, 3},
		{"Case 3", args{[]int{1}}, 1},
		{"Case 3", args{[]int{2, 2, 2, 5, 9, 10}}, 1},
		{"Case 3", args{[]int{2, 2, 2, 8, 8, 8}}, 2},
		{"Case 3", args{[]int{2, 2, 2, 5, 8, 8, 8}}, 1},
		{"Case 3", args{[]int{4, 4, 4, 7, 8, 8, 8}}, 1},
		{"Case 3", args{[]int{3, 3, 3, 9, 9, 9}}, 2},
		{"Case 3", args{[]int{9, 9, 9, 11, 11, 11}}, 2},
		{"Case 3", args{[]int{9, 9, 9, 9, 9, 9, 11, 11, 11}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumArrayLength(tt.args.nums); got != tt.want {
				t.Errorf("minimumArrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
