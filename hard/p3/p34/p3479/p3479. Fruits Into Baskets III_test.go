package p3479

import "testing"

func Test_numOfUnplacedFruits(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{fruits: []int{4, 2, 5}, baskets: []int{3, 5, 4}}, 1},
		{"Example 2", args{fruits: []int{3, 6, 1}, baskets: []int{6, 4, 7}}, 0},
		{"Testcase 45", args{fruits: []int{26}, baskets: []int{17}}, 1},
		{"Testcase 45", args{fruits: []int{35, 61}, baskets: []int{76, 56}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfUnplacedFruits(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("numOfUnplacedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
