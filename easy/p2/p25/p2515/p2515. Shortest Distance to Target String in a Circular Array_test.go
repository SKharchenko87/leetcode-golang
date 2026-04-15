package p2515

import "testing"

func Test_closestTarget(t *testing.T) {
	type args struct {
		words      []string
		target     string
		startIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{words: []string{"hello", "i", "am", "leetcode", "hello"}, target: "hello", startIndex: 1}, 1},
		{"Example 2", args{words: []string{"a", "b", "leetcode"}, target: "leetcode", startIndex: 0}, 1},
		{"Example 3", args{words: []string{"i", "eat", "leetcode"}, target: "ate", startIndex: 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestTarget(tt.args.words, tt.args.target, tt.args.startIndex); got != tt.want {
				t.Errorf("closestTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
