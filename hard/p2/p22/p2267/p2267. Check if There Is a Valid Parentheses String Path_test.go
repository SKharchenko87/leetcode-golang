package p2267

import "testing"

func Test_hasValidPath(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1:", args{grid: [][]byte{{'(', '(', '('}, {')', '(', ')'}, {'(', '(', ')'}, {'(', '(', ')'}}}, true},
		{"Example 2:", args{grid: [][]byte{{')', ')'}, {'(', '('}}}, false},
		{"TestCase 54:", args{grid: [][]byte{{'(', ')'}, {'(', ')'}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasValidPath(tt.args.grid); got != tt.want {
				t.Errorf("hasValidPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
