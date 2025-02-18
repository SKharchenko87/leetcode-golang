package p2698

import "testing"

func Test_punishmentNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 10}, 182},
		{"Example 2", args{n: 37}, 1478},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := punishmentNumber(tt.args.n); got != tt.want {
				t.Errorf("punishmentNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
