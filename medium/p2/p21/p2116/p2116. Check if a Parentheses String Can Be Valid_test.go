package p2116

import "testing"

func Test_canBeValid(t *testing.T) {
	type args struct {
		s      string
		locked string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s: "))()))", locked: "010100"}, true},
		{"Example 2", args{s: "()()", locked: "0000"}, true},
		{"Example 3", args{s: ")", locked: "0"}, false},
		{"Example 11", args{s: "()", locked: "11"}, true},
		{"Example 207", args{s: "())(()(()(())()())(())((())(()())((())))))(((((((())(()))))(", locked: "100011110110011011010111100111011101111110000101001101001111"}, false},
		{"Example 256", args{s: "(())()", locked: "110110"}, true},
		{"Example 256", args{s: "(((())", locked: "011100"}, false},
		{"Test", args{s: "()()()", locked: "110110"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeValid(tt.args.s, tt.args.locked); got != tt.want {
				t.Errorf("canBeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ((*)(*

// ((*)(*

// (*(*
