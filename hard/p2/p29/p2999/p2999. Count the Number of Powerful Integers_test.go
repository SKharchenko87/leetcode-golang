package p2999

import "testing"

func Test_numberOfPowerfulInt(t *testing.T) {
	type args struct {
		start  int64
		finish int64
		limit  int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{start: 1, finish: 6000, limit: 4, s: "124"}, 5},
		{"Example 2", args{start: 15, finish: 215, limit: 6, s: "10"}, 2},
		{"Example 3", args{start: 1000, finish: 2000, limit: 4, s: "3000"}, 0},
		{"My 1", args{start: 1, finish: 1100, limit: 4, s: "44"}, 6},
		{"My 2", args{start: 1, finish: 1100, limit: 4, s: "4"}, 30},
		{"My 3", args{start: 1, finish: 6000, limit: 4, s: "24"}, 25},
		{"My 4", args{start: 1, finish: 6000, limit: 4, s: "4"}, 125},
		{"My 5", args{start: 1, finish: 6660, limit: 4, s: "4"}, 125},
		{"My 6", args{start: 1, finish: 4440, limit: 4, s: "4"}, 124},
		{"My 6", args{start: 1, finish: 4444, limit: 4, s: "4"}, 125},
		{"TestCase 447", args{start: 1, finish: 971, limit: 9, s: "72"}, 9},
		{"TestCase 545", args{start: 20, finish: 1159, limit: 5, s: "20"}, 8},
		{"TestCase 551", args{start: 1114, finish: 1864854501, limit: 7, s: "26"}, 4194295},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPowerfulInt(tt.args.start, tt.args.finish, tt.args.limit, tt.args.s); got != tt.want {
				t.Errorf("numberOfPowerfulInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
