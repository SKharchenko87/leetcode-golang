package p1318

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{2, 6, 5}, 3},
		{"Case 1", args{4, 2, 7}, 1},
		{"Case 1", args{1, 2, 3}, 0},
		{"Case 1", args{10, 9, 1}, 3},
		{"Case 1", args{7, 3, 9}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
