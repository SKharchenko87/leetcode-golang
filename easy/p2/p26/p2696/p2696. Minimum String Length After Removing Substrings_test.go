package p2696

import "testing"

func Test_minLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "ABFCACDB"}, 2},
		{"Example 2", args{s: "ACBBD"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minLength(tt.args.s); got != tt.want {
				t.Errorf("minLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
