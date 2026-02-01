package p0744

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Example 1", args{letters: []byte{'c', 'f', 'j'}, target: 'a'}, 'c'},
		{"Example 2", args{letters: []byte{'c', 'f', 'j'}, target: 'c'}, 'f'},
		{"Example 3", args{letters: []byte{'x', 'x', 'y', 'y'}, target: 'z'}, 'x'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
