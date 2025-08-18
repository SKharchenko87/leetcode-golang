package p0679

import "testing"

func Test_judgePoint24(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{cards: []int{4, 1, 8, 7}}, true},
		{"Example 2", args{cards: []int{1, 2, 1, 2}}, false},
		{"TestCase 63", args{cards: []int{1, 5, 9, 1}}, false},
		{"TestCase 70", args{cards: []int{3, 4, 6, 7}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgePoint24(tt.args.cards); got != tt.want {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}
