package p0135

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{ratings: []int{1, 0, 2}}, 5},
		{"Example 2", args{ratings: []int{1, 2, 2}}, 4},
		{"TestCase 10", args{ratings: []int{1, 2, 87, 87, 87, 2, 1}}, 13},
		{"TestCase 45", args{ratings: []int{1, 2, 4, 4, 3}}, 9},
		{"TestCase 48", args{ratings: []int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
