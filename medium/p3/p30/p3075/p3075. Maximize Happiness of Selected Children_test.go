package p3075

import "testing"

func Test_maximumHappinessSum(t *testing.T) {
	type args struct {
		happiness []int
		k         int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{[]int{1, 2, 3}, 2}, 4},
		{"Example 2", args{[]int{1, 1, 1, 1}, 2}, 1},
		{"Example 3", args{[]int{2, 3, 4, 5}, 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumHappinessSum(tt.args.happiness, tt.args.k); got != tt.want {
				t.Errorf("maximumHappinessSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
