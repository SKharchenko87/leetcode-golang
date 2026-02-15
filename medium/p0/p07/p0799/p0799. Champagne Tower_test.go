package p0799

import "testing"

func Test_champagneTower(t *testing.T) {
	type args struct {
		poured      int
		query_row   int
		query_glass int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{poured: 1, query_row: 1, query_glass: 1}, 0.00000},
		{"Example 2", args{poured: 2, query_row: 1, query_glass: 1}, 0.50000},
		{"Example 3", args{poured: 100000009, query_row: 33, query_glass: 17}, 1.00000},
		{"My 1", args{poured: 7, query_row: 3, query_glass: 2}, 0.5},
		{"My 2", args{poured: 10, query_row: 4, query_glass: 2}, 0.625},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := champagneTower(tt.args.poured, tt.args.query_row, tt.args.query_glass); got != tt.want {
				t.Errorf("champagneTower() = %v, want %v", got, tt.want)
			}
		})
	}
}
