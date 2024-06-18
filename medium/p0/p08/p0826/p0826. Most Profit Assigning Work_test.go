package p0826

import "testing"

func Test_maxProfitAssignment(t *testing.T) {
	type args struct {
		difficulty []int
		profit     []int
		worker     []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"Example 1", args{difficulty: []int{2, 4, 6, 8, 10}, profit: []int{10, 20, 30, 40, 50}, worker: []int{4, 5, 6, 7}}, 100},
		{"Example 2", args{difficulty: []int{85, 47, 57}, profit: []int{24, 66, 99}, worker: []int{40, 25, 25}}, 0},
		{"Test 17", args{difficulty: []int{5, 50, 92, 21, 24, 70, 17, 63, 30, 53}, profit: []int{68, 100, 3, 99, 56, 43, 26, 93, 55, 25}, worker: []int{96, 3, 55, 30, 11, 58, 68, 36, 26, 1}}, 765},
		{"Test 51", args{difficulty: []int{68, 35, 52, 47, 86}, profit: []int{67, 17, 1, 81, 3}, worker: []int{92, 10, 85, 84, 82}}, 324},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitAssignment(tt.args.difficulty, tt.args.profit, tt.args.worker); got != tt.want {
				t.Errorf("maxProfitAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
