package p1780

import "testing"

func Test_checkPowersOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 12}, true},
		{"Example 2", args{n: 91}, true},
		{"Example 3", args{n: 21}, false},
		{"Example 3", args{n: 4782970}, true},
		{"Example 3", args{n: 4782968}, false},
		{"Example 3", args{n: 85}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPowersOfThree(tt.args.n); got != tt.want {
				t.Errorf("checkPowersOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
