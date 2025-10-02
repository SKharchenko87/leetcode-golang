package p3307

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k          int64
		operations []int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Example 1", args{k: 5, operations: []int{0, 0, 0}}, 'a'},
		{"Example 2", args{k: 10, operations: []int{0, 1, 0, 1}}, 'b'},
		{"My 0", args{k: 16, operations: []int{0, 1, 0, 1}}, 'c'},
		{"My 1", args{k: 13, operations: []int{0, 1, 0, 1}}, 'b'},
		{"TestCase 2", args{k: 1, operations: []int{0}}, 'a'},
		{"TestCase 14", args{k: 1, operations: []int{0, 1}}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k, tt.args.operations); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
