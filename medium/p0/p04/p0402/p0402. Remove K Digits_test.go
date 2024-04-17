package p0402

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"1432219", 3}, "1219"},
		{"Example 2", args{"10200", 1}, "200"},
		{"Example 3", args{"10", 2}, "0"},
		{"Example 3", args{"10", 1}, "0"},
		{"Example 3", args{"112", 1}, "11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
