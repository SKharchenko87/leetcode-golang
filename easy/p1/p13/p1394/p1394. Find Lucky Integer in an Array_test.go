package p1394

import "testing"

func Test_findLucky(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arr: []int{2, 2, 3, 4}}, 2},
		{"Example 2", args{arr: []int{1, 2, 2, 3, 3, 3}}, 3},
		{"Example 3", args{arr: []int{2, 2, 2, 3, 3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLucky(tt.args.arr); got != tt.want {
				t.Errorf("findLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
