package p3272

import "testing"

func Test_countGoodIntegers(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{n: 3, k: 5}, 27},
		{"Example 2", args{n: 1, k: 4}, 2},
		{"Example 3", args{n: 5, k: 6}, 2468},
		{"TestCase 54", args{n: 4, k: 1}, 252},
		{"TestCase 81", args{n: 10, k: 1}, 41457024},
		{"TestCase n=1	k=1", args{n: 1, k: 1}, 9},
		{"TestCase n=1	k=2", args{n: 1, k: 2}, 4},
		{"TestCase n=1	k=3", args{n: 1, k: 3}, 3},
		{"TestCase n=1	k=4", args{n: 1, k: 4}, 2},
		{"TestCase n=1	k=5", args{n: 1, k: 5}, 1},
		{"TestCase n=1	k=6", args{n: 1, k: 6}, 1},
		{"TestCase n=1	k=7", args{n: 1, k: 7}, 1},
		{"TestCase n=1	k=8", args{n: 1, k: 8}, 1},
		{"TestCase n=1	k=9", args{n: 1, k: 9}, 1},
		{"TestCase n=2	k=1", args{n: 2, k: 1}, 9},
		{"TestCase n=2	k=2", args{n: 2, k: 2}, 4},
		{"TestCase n=2	k=3", args{n: 2, k: 3}, 3},
		{"TestCase n=2	k=4", args{n: 2, k: 4}, 2},
		{"TestCase n=2	k=5", args{n: 2, k: 5}, 1},
		{"TestCase n=2	k=6", args{n: 2, k: 6}, 1},
		{"TestCase n=2	k=7", args{n: 2, k: 7}, 1},
		{"TestCase n=2	k=8", args{n: 2, k: 8}, 1},
		{"TestCase n=2	k=9", args{n: 2, k: 9}, 1},
		{"TestCase n=3	k=1", args{n: 3, k: 1}, 243},
		{"TestCase n=3	k=2", args{n: 3, k: 2}, 108},
		{"TestCase n=3	k=3", args{n: 3, k: 3}, 69},
		{"TestCase n=3	k=4", args{n: 3, k: 4}, 54},
		{"TestCase n=3	k=5", args{n: 3, k: 5}, 27},
		{"TestCase n=3	k=6", args{n: 3, k: 6}, 30},
		{"TestCase n=3	k=7", args{n: 3, k: 7}, 33},
		{"TestCase n=3	k=8", args{n: 3, k: 8}, 27},
		{"TestCase n=3	k=9", args{n: 3, k: 9}, 23},
		{"TestCase n=4	k=1", args{n: 4, k: 1}, 252},
		{"TestCase n=4	k=2", args{n: 4, k: 2}, 172},
		{"TestCase n=4	k=3", args{n: 4, k: 3}, 84},
		{"TestCase n=4	k=4", args{n: 4, k: 4}, 98},
		{"TestCase n=4	k=5", args{n: 4, k: 5}, 52},
		{"TestCase n=4	k=6", args{n: 4, k: 6}, 58},
		{"TestCase n=4	k=7", args{n: 4, k: 7}, 76},
		{"TestCase n=4	k=8", args{n: 4, k: 8}, 52},
		{"TestCase n=4	k=9", args{n: 4, k: 9}, 28},
		{"TestCase n=5	k=1", args{n: 5, k: 1}, 10935},
		{"TestCase n=5	k=2", args{n: 5, k: 2}, 7400},
		{"TestCase n=5	k=3", args{n: 5, k: 3}, 3573},
		{"TestCase n=5	k=4", args{n: 5, k: 4}, 4208},
		{"TestCase n=5	k=5", args{n: 5, k: 5}, 2231},
		{"TestCase n=5	k=6", args{n: 5, k: 6}, 2468},
		{"TestCase n=5	k=7", args{n: 5, k: 7}, 2665},
		{"TestCase n=5	k=8", args{n: 5, k: 8}, 2231},
		{"TestCase n=5	k=9", args{n: 5, k: 9}, 1191},
		{"TestCase n=6	k=1", args{n: 6, k: 1}, 10944},
		{"TestCase n=6	k=2", args{n: 6, k: 2}, 9064},
		{"TestCase n=6	k=3", args{n: 6, k: 3}, 3744},
		{"TestCase n=6	k=4", args{n: 6, k: 4}, 6992},
		{"TestCase n=6	k=5", args{n: 6, k: 5}, 3256},
		{"TestCase n=6	k=6", args{n: 6, k: 6}, 3109},
		{"TestCase n=6	k=7", args{n: 6, k: 7}, 3044},
		{"TestCase n=6	k=8", args{n: 6, k: 8}, 5221},
		{"TestCase n=6	k=9", args{n: 6, k: 9}, 1248},
		{"TestCase n=7	k=1", args{n: 7, k: 1}, 617463},
		{"TestCase n=7	k=2", args{n: 7, k: 2}, 509248},
		{"TestCase n=7	k=3", args{n: 7, k: 3}, 206217},
		{"TestCase n=7	k=4", args{n: 7, k: 4}, 393948},
		{"TestCase n=7	k=5", args{n: 7, k: 5}, 182335},
		{"TestCase n=7	k=6", args{n: 7, k: 6}, 170176},
		{"TestCase n=7	k=7", args{n: 7, k: 7}, 377610},
		{"TestCase n=7	k=8", args{n: 7, k: 8}, 292692},
		{"TestCase n=7	k=9", args{n: 7, k: 9}, 68739},
		{"TestCase n=8	k=1", args{n: 8, k: 1}, 617472},
		{"TestCase n=8	k=2", args{n: 8, k: 2}, 563392},
		{"TestCase n=8	k=3", args{n: 8, k: 3}, 207840},
		{"TestCase n=8	k=4", args{n: 8, k: 4}, 494818},
		{"TestCase n=8	k=5", args{n: 8, k: 5}, 237112},
		{"TestCase n=8	k=6", args{n: 8, k: 6}, 188945},
		{"TestCase n=8	k=7", args{n: 8, k: 7}, 506388},
		{"TestCase n=8	k=8", args{n: 8, k: 8}, 460048},
		{"TestCase n=8	k=9", args{n: 8, k: 9}, 69280},
		{"TestCase n=9	k=1", args{n: 9, k: 1}, 41457015},
		{"TestCase n=9	k=2", args{n: 9, k: 2}, 37728000},
		{"TestCase n=9	k=3", args{n: 9, k: 3}, 13726509},
		{"TestCase n=9	k=4", args{n: 9, k: 4}, 33175696},
		{"TestCase n=9	k=5", args{n: 9, k: 5}, 15814071},
		{"TestCase n=9	k=6", args{n: 9, k: 6}, 12476696},
		{"TestCase n=9	k=7", args{n: 9, k: 7}, 36789447},
		{"TestCase n=9	k=8", args{n: 9, k: 8}, 30771543},
		{"TestCase n=9	k=9", args{n: 9, k: 9}, 4623119},
		{"TestCase n=10	k=1", args{n: 10, k: 1}, 41457024},
		{"TestCase n=10	k=2", args{n: 10, k: 2}, 39718144},
		{"TestCase n=10	k=3", args{n: 10, k: 3}, 13831104},
		{"TestCase n=10	k=4", args{n: 10, k: 4}, 37326452},
		{"TestCase n=10	k=5", args{n: 10, k: 5}, 19284856},
		{"TestCase n=10	k=6", args{n: 10, k: 6}, 13249798},
		{"TestCase n=10	k=7", args{n: 10, k: 7}, 40242031},
		{"TestCase n=10	k=8", args{n: 10, k: 8}, 35755906},
		{"TestCase n=10	k=9", args{n: 10, k: 9}, 4610368},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodIntegers(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("countGoodIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDigits(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
	}{
		{"Test 1", args{num: 20202}, 22200, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getDigits(tt.args.num)
			if got != tt.want {
				t.Errorf("getDigits() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getDigits() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
