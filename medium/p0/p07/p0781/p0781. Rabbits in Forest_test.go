package p0781

import "testing"

func Test_numRabbits(t *testing.T) {
	type args struct {
		answers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{answers: []int{1, 1, 2}}, 5},
		{"Example 2", args{answers: []int{10, 10, 10}}, 11},
		{"My 1", args{answers: []int{10, 10, 10, 10, 10, 10, 10, 10, 10}}, 11},
		{"My 2", args{answers: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}}, 11},
		{"My 3", args{answers: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}}, 11},
		{"My 4", args{answers: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}}, 22},
		{"My 4", args{answers: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}}, 22},
		{"TestCase 21", args{answers: []int{1, 0, 1, 0, 0}}, 5},
		{"TestCase 51", args{answers: []int{0, 1, 0, 2, 0, 1, 0, 2, 1, 1}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.args.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
