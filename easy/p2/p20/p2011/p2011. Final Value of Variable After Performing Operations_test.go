package p2011

import "testing"

func Test_finalValueAfterOperations(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{operations: []string{"--X", "X++", "X++"}}, 1},
		{"Example 2", args{operations: []string{"++X", "++X", "X++"}}, 3},
		{"Example 3", args{operations: []string{"X++", "++X", "--X", "X--"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalValueAfterOperations(tt.args.operations); got != tt.want {
				t.Errorf("finalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
