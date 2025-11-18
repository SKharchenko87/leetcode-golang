package p0717

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{bits: []int{1, 0, 0}}, true},
		{"Example 2", args{bits: []int{1, 1, 1, 0}}, false},
		{"My 0,0", args{[]int{0, 0}}, true},
		{"My 1,0", args{[]int{1, 0}}, false},
		{"My 0,0,0", args{[]int{0, 0, 0}}, true},
		{"My 0,1,0", args{[]int{0, 1, 0}}, false},
		{"My 1,0,0", args{[]int{1, 0, 0}}, true},
		{"My 1,1,0", args{[]int{1, 1, 0}}, true},
		{"My 0,0,0,0", args{[]int{0, 0, 0, 0}}, true},
		{"My 0,0,1,0", args{[]int{0, 0, 1, 0}}, false},
		{"My 0,1,0,0", args{[]int{0, 1, 0, 0}}, true},
		{"My 0,1,1,0", args{[]int{0, 1, 1, 0}}, true},
		{"My 1,0,0,0", args{[]int{1, 0, 0, 0}}, true},
		{"My 1,0,1,0", args{[]int{1, 0, 1, 0}}, false},
		{"My 1,1,0,0", args{[]int{1, 1, 0, 0}}, true},
		{"My 1,1,1,0", args{[]int{1, 1, 1, 0}}, false},
		{"My 0,0,0,0,0", args{[]int{0, 0, 0, 0, 0}}, true},
		{"My 0,0,0,1,0", args{[]int{0, 0, 0, 1, 0}}, false},
		{"My 0,0,1,0,0", args{[]int{0, 0, 1, 0, 0}}, true},
		{"My 0,0,1,1,0", args{[]int{0, 0, 1, 1, 0}}, true},
		{"My 0,1,0,0,0", args{[]int{0, 1, 0, 0, 0}}, true},
		{"My 0,1,0,1,0", args{[]int{0, 1, 0, 1, 0}}, false},
		{"My 0,1,1,0,0", args{[]int{0, 1, 1, 0, 0}}, true},
		{"My 0,1,1,1,0", args{[]int{0, 1, 1, 1, 0}}, false},
		{"My 1,0,0,0,0", args{[]int{1, 0, 0, 0, 0}}, true},
		{"My 1,0,0,1,0", args{[]int{1, 0, 0, 1, 0}}, false},
		{"My 1,0,1,0,0", args{[]int{1, 0, 1, 0, 0}}, true},
		{"My 1,0,1,1,0", args{[]int{1, 0, 1, 1, 0}}, true},
		{"My 1,1,0,0,0", args{[]int{1, 1, 0, 0, 0}}, true},
		{"My 1,1,0,1,0", args{[]int{1, 1, 0, 1, 0}}, false},
		{"My 1,1,1,0,0", args{[]int{1, 1, 1, 0, 0}}, true},
		{"My 1,1,1,1,0", args{[]int{1, 1, 1, 1, 0}}, true},
		{"My 1,1,1,1,1,0", args{[]int{1, 1, 1, 1, 1, 0}}, false},
		{"My 0,1,1,1,1,0", args{[]int{0, 1, 1, 1, 1, 0}}, true},
		{"My 1,0,1,1,1,0", args{[]int{1, 0, 1, 1, 1, 0}}, false},
		{"My 1,0,0,1,1,0", args{[]int{1, 0, 0, 1, 1, 0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneBitCharacter(tt.args.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
