package p0621

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2}, 8},
		{"Example 2", args{[]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1}, 6},
		{"Example 3", args{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3}, 10},
		{"Example 60", args{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
