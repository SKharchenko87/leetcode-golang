package p2429

import "testing"

func Test_minimizeXor(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{num1: 3, num2: 5}, 3},
		{"Example 2", args{num1: 1, num2: 12}, 3},
		{"Testcase 164", args{num1: 39446869, num2: 13801288}, 39446848},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeXor(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("minimizeXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
