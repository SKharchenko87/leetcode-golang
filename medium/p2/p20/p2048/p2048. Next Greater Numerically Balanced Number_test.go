package p2048

import "testing"

func Test_nextBeautifulNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 1}, 22},
		{"Example 2", args{n: 1000}, 1333},
		{"Example 3", args{n: 3000}, 3133},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextBeautifulNumber(tt.args.n); got != tt.want {
				t.Errorf("nextBeautifulNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
