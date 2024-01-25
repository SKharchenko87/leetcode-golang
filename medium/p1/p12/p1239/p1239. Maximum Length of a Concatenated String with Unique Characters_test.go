package p1239

import "testing"

func Test_maxLength(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]string{"un", "iq", "ue"}}, 4},
		{"Case 2", args{[]string{"cha", "r", "act", "ers"}}, 6},
		{"Case 3", args{[]string{"abcdefghijklmnopqrstuvwxyz"}}, 26},
		{"Case 3", args{[]string{"aa", "bb"}}, 0},
		{"My 1", args{[]string{"a"}}, 1},
		{"My 1", args{[]string{"ab"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.args.arr); got != tt.want {
				t.Errorf("maxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
