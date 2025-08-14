package p0552

import "testing"

func Test_checkRecord(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{2}, 8},
		{"Example 2", args{1}, 3},
		{"Example 3", args{10101}, 183236316},
		{"*Example 28", args{28}, 530803311},
		{"*Example 28", args{29}, 9569297},
		{"*Example 3", args{3}, 19},
		{"*Example 4", args{4}, 43},
		{"*Example 5", args{5}, 94},
		{"*Example 5", args{6}, 200},
		{"*Example 5", args{7}, 418},
		{"*Example 5", args{8}, 861},
		{"*Example 99991", args{99991}, 991321922},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord(tt.args.n); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
