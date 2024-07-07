package p2582

import "testing"

func Test_passThePillow(t *testing.T) {
	type args struct {
		n    int
		time int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 4, time: 5}, 2},
		{"Example 2", args{n: 3, time: 2}, 3},
		{"Example 47", args{n: 18, time: 38}, 5},
		{"Example 47", args{n: 6, time: 14}, 5},
		{"Example 47", args{n: 6, time: 20}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := passThePillow(tt.args.n, tt.args.time); got != tt.want {
				t.Errorf("passThePillow() = %v, want %v", got, tt.want)
			}
		})
	}
}
