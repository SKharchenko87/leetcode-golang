package p2075

import "testing"

func Test_decodeCiphertext(t *testing.T) {
	type args struct {
		encodedText string
		rows        int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{encodedText: "ch   ie   pr", rows: 3}, "cipher"},
		{"Example 2", args{encodedText: "iveo    eed   l te   olc", rows: 4}, "i love leetcode"},
		{"Example 3", args{encodedText: "coding", rows: 1}, "coding"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeCiphertext(tt.args.encodedText, tt.args.rows); got != tt.want {
				t.Errorf("decodeCiphertext() = %v, want %v", got, tt.want)
			}
		})
	}
}
