package p0273

import "testing"

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{num: 123}, "One Hundred Twenty Three"},
		{"Example 2", args{num: 12345}, "Twelve Thousand Three Hundred Forty Five"},
		{"Example 3", args{num: 1234567}, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{"Example 3", args{num: 20}, "Twenty"},
		{"Example 594", args{num: 3200348}, "Three Million Two Hundred Thousand Three Hundred Forty Eight"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
