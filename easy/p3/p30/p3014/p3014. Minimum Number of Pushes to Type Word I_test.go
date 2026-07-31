package p3014

import "testing"

func Test_minimumPushes(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{word: "abcde"}, 5},
		{"Example 2", args{word: "xycdefghij"}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPushes(tt.args.word); got != tt.want {
				t.Errorf("minimumPushes() = %v, want %v", got, tt.want)
			}
		})
	}
}
