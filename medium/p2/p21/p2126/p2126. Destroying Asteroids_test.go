package p2126

import "testing"

func Test_asteroidsDestroyed(t *testing.T) {
	type args struct {
		mass      int
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{mass: 10, asteroids: []int{3, 9, 19, 5, 21}}, true},
		{"Example 2", args{mass: 5, asteroids: []int{4, 9, 23, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidsDestroyed(tt.args.mass, tt.args.asteroids); got != tt.want {
				t.Errorf("asteroidsDestroyed() = %v, want %v", got, tt.want)
			}
		})
	}
}
