package p3227

import "testing"

func Test_doesAliceWin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s: "leetcoder"}, true},
		{"Example 2", args{s: "bbcd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesAliceWin(tt.args.s); got != tt.want {
				t.Errorf("doesAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
