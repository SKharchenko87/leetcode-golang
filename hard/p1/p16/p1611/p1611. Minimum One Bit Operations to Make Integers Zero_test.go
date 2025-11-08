package p1611

import "testing"

func Test_minimumOneBitOperations(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 3}, 2},
		{"Example 2", args{n: 6}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOneBitOperations(tt.args.n); got != tt.want {
				t.Errorf("minimumOneBitOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
