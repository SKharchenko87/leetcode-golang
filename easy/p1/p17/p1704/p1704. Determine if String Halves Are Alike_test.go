package p1704

import "testing"

func Test_halvesAreAlike(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{"book"}, true},
		{"Case 2", args{"textbook"}, false},
		{"Case 57", args{"AbCdEfGh"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halvesAreAlike(tt.args.s); got != tt.want {
				t.Errorf("halvesAreAlike() = %v, want %v", got, tt.want)
			}
		})
	}
}
