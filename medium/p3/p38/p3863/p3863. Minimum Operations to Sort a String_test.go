package p3863

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "dog"}, 1},
		{"Example 2", args{s: "card"}, 2},
		{"Example 3", args{s: "gf"}, -1},
		{"TestCase 1170", args{s: "jgg"}, 2},
		{"TestCase 1219", args{s: "okmk"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.s); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
