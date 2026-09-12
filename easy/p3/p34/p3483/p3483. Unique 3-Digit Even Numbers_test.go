package p3483

import "testing"

func Test_totalNumbers(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{digits: []int{1, 2, 3, 4}}, 12},
		{"Example 2", args{digits: []int{0, 2, 2}}, 2},
		{"Example 3", args{digits: []int{6, 6, 6}}, 1},
		{"Example 4", args{digits: []int{1, 3, 5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNumbers(tt.args.digits); got != tt.want {
				t.Errorf("totalNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
