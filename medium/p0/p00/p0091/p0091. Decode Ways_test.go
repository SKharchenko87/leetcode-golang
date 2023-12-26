package p0091

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{"12"}, 2},
		{"Case 2", args{"226"}, 3},
		{"Case 3", args{"06"}, 0},
		{"Case 4", args{"1201234"}, 3},
		{"Case 5", args{"2611055971756562"}, 4},
		{"Case 6", args{"111111111111111111111111111111111111111111111"}, 1836311903},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
