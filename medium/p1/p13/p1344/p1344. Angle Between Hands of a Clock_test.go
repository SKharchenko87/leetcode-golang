package p1344

import "testing"

func Test_angleClock(t *testing.T) {
	type args struct {
		hour    int
		minutes int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{hour: 12, minutes: 30}, 165},
		{"Example 2", args{hour: 3, minutes: 30}, 75},
		{"Example 3", args{hour: 3, minutes: 15}, 7.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angleClock(tt.args.hour, tt.args.minutes); got != tt.want {
				t.Errorf("angleClock() = %v, want %v", got, tt.want)
			}
		})
	}
}
