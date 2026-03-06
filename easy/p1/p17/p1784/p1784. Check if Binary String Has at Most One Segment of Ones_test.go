package p1784

import "testing"

func Test_checkOnesSegment(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s: "1001"}, false},
		{"Example 2", args{s: "110"}, true},
		{"My 1", args{s: "11011"}, false},
		{"My 2", args{s: "110101010101010"}, false},
		{"My 3", args{s: "101010101101010"}, false},
		{"My 4", args{s: "101010101010100"}, false},
		{"My 5", args{s: "101010101001011"}, false},
		{"My 6", args{s: "1"}, true},
		{"My 7", args{s: "10"}, true},
		{"My 8", args{s: "11"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkOnesSegment(tt.args.s); got != tt.want {
				t.Errorf("checkOnesSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
