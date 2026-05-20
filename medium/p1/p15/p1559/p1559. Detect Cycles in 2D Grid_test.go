package p1559

import "testing"

func Test_containsCycle(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{{"Example 1", args{grid: [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}}, true},
		{"Example 2", args{grid: [][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}}}, true},
		{"Example 3", args{grid: [][]byte{{'a', 'b', 'b'}, {'b', 'z', 'b'}, {'b', 'b', 'a'}}}, false},
		{"My 1", args{grid: [][]byte{{'b', 'a', 'b', 'a', 'b'}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsCycle(tt.args.grid); got != tt.want {
				t.Errorf("containsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
