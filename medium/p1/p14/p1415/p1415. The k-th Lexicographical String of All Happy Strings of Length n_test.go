package p1415

import "testing"

func Test_getHappyString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{n: 1, k: 3}, "c"},
		{"Example 2", args{n: 1, k: 4}, ""},
		{"Example 3", args{n: 3, k: 9}, "cab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHappyString(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getHappyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
