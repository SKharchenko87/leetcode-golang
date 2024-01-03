package p2125

import "testing"

func Test_numberOfBeams(t *testing.T) {
	type args struct {
		bank []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 0", args{[]string{"000", "000", "000"}}, 0},
		{"Case 0", args{[]string{"000", "000", "111"}}, 0},
		{"Case 1", args{[]string{"011001", "000000", "010100", "001000"}}, 8},
		{"Case 2", args{[]string{"000", "111", "000"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBeams(tt.args.bank); got != tt.want {
				t.Errorf("numberOfBeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
