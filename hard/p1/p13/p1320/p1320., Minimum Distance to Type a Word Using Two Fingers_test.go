package p1320

import "testing"

func Test_minimumDistance(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{word: "CAKE"}, 3},
		{"Example 2", args{word: "HAPPY"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDistance(tt.args.word); got != tt.want {
				t.Errorf("minimumDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
