package p3838

import "testing"

func Test_mapWordWeights(t *testing.T) {
	type args struct {
		words   []string
		weights []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{words: []string{"abcd", "def", "xyz"}, weights: []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}}, "rij"},
		{"Example 2", args{words: []string{"a", "b", "c"}, weights: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}, "yyy"},
		{"Example 3", args{words: []string{"abcd"}, weights: []int{7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5}}, "g"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapWordWeights(tt.args.words, tt.args.weights); got != tt.want {
				t.Errorf("mapWordWeights() = %v, want %v", got, tt.want)
			}
		})
	}
}
