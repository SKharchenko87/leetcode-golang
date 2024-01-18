package p0790

import "testing"

func Test_numTilings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{1}, 1},
		{"Case 2", args{2}, 2},
		{"Case 3", args{3}, 5},
		{"Case 4", args{4}, 11},
		{"Case 5", args{5}, 24},
		{"Case 6", args{6}, 53},
		{"Case 7", args{7}, 117},
		{"Case 8", args{8}, 258},
		{"Case 14", args{57}, 643607090},
		{"Case 14", args{60}, 882347204},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilings(tt.args.n); got != tt.want {
				t.Errorf("numTilings() = %v, want %v", got, tt.want)
			}
		})
	}
}
