package p3136

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{word: "234Adas"}, true},
		{"Example 2", args{word: "b3"}, false},
		{"Example 3", args{word: "a3$e"}, false},
		{"TestCase 495", args{word: "234Adas"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.word); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
