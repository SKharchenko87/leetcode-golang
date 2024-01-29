package p0629

import "testing"

func Test_kInversePairs(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{3, 0}, 1},
		{"Case 2", args{3, 1}, 2},
		{"Case 3", args{7, 6}, 259},
		{"Case 4", args{7, 7}, 359},
		{"Case 5", args{7, 8}, 455},
		{"Case 6", args{7, 9}, 531},
		{"Case 7", args{7, 10}, 573},
		{"Case 8", args{7, 11}, 573},
		{"Case 1000", args{1000, 1000}, 663677020},
		{"Case 1000", args{1000, 999}, 40778721},
		{"Case 55_51", args{55, 51}, 783257836},
		{"Case 45_67", args{45, 67}, 668985185},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kInversePairs(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kInversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
