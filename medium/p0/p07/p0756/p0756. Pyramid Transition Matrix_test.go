package p0756

import "testing"

func Test_pyramidTransition(t *testing.T) {
	type args struct {
		bottom  string
		allowed []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{bottom: "BCD", allowed: []string{"BCC", "CDE", "CEA", "FFF"}}, true},
		{"Example 2", args{bottom: "AAAA", allowed: []string{"AAB", "AAC", "BCD", "BBE", "DEF"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pyramidTransition(tt.args.bottom, tt.args.allowed); got != tt.want {
				t.Errorf("pyramidTransition() = %v, want %v", got, tt.want)
			}
		})
	}
}
