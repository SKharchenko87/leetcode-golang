package p3461

import "testing"

func Test_hasSameDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s: "3902"}, true},
		{"Example 2", args{s: "34789"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSameDigits(tt.args.s); got != tt.want {
				t.Errorf("hasSameDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
