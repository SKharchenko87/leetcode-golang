package p1052

import "testing"

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{customers: []int{1, 0, 1, 2, 1, 1, 7, 5}, grumpy: []int{0, 1, 0, 1, 0, 1, 0, 1}, minutes: 3}, 16},
		{"Example 2", args{customers: []int{1}, grumpy: []int{0}, minutes: 1}, 1},
		{"Testcase 61", args{customers: []int{3}, grumpy: []int{1}, minutes: 1}, 3},
		{"Testcase 73", args{customers: []int{4, 10, 10}, grumpy: []int{1, 1, 0}, minutes: 2}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.minutes); got != tt.want {
				t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
