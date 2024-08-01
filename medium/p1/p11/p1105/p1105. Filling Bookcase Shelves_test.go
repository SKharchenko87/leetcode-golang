package p1105

import "testing"

func Test_minHeightShelves(t *testing.T) {
	type args struct {
		books      [][]int
		shelfWidth int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{books: [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, shelfWidth: 4}, 6},
		{"Example 2", args{books: [][]int{{1, 3}, {2, 4}, {3, 2}}, shelfWidth: 6}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minHeightShelves(tt.args.books, tt.args.shelfWidth); got != tt.want {
				t.Errorf("minHeightShelves() = %v, want %v", got, tt.want)
			}
		})
	}
}
