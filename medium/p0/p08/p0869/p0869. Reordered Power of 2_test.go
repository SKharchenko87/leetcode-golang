package p0869

import "testing"

func Test_reorderedPowerOf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 1}, true},
		{"Example 2", args{n: 10}, false},
		{"My 4210", args{n: 4210}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.args.n); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
