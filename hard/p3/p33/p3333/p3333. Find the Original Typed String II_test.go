package p3333

import "testing"

func Test_possibleStringCount(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{word: "aabbccdd", k: 7}, 5},
		{"Example 2", args{word: "aabbccdd", k: 8}, 1},
		{"Example 3", args{word: "aaabbb", k: 3}, 8},
		{"TestCase 782", args{word: "aaafffffeeeeoooooobbbbbbbbggghhhhhhheeeeeeewwwwwwwwiiiiiiiijjjjjjjjaaaannnnnnniiiiiiimmmmmmmmrrrrrrrriieeeiiiiiiyyyybbbbbbbbbbbbbbbbbbbbzzzzeelllllllliiippppmmmmmmmmmmuuuuuuunnnnnnnnzzzzzzzzzzijjjjvvb", k: 118}, 297466760},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleStringCount(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("possibleStringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
