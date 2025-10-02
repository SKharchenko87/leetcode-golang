package p2071

import "testing"

func Test_maxTaskAssign(t *testing.T) {
	type args struct {
		tasks    []int
		workers  []int
		pills    int
		strength int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{tasks: []int{3, 2, 1}, workers: []int{0, 3, 3}, pills: 1, strength: 1}, 3},
		{"Example 2", args{tasks: []int{5, 4}, workers: []int{0, 0, 0}, pills: 1, strength: 5}, 1},
		{"Example 3", args{tasks: []int{10, 15, 30}, workers: []int{0, 10, 10, 10, 10}, pills: 3, strength: 10}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTaskAssign(tt.args.tasks, tt.args.workers, tt.args.pills, tt.args.strength); got != tt.want {
				t.Errorf("maxTaskAssign() = %v, want %v", got, tt.want)
			}
		})
	}
}
