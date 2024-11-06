package p3011

import "testing"

func Test_canSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{8, 4, 2, 30, 15}}, true},
		{"Example 2", args{nums: []int{1, 2, 3, 4, 5}}, true},
		{"Example 3", args{nums: []int{3, 16, 8, 4, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSortArray(tt.args.nums); got != tt.want {
				t.Errorf("canSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOneBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"Test 1", args{num: 0}, 0},
		{"Test 1", args{num: 1}, 1},
		{"Test 1", args{num: 2}, 1},
		{"Test 1", args{num: 3}, 2},
		{"Test 1", args{num: 4}, 1},
		{"Test 1", args{num: 255}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countOneBits(tt.args.num); gotRes != tt.wantRes {
				t.Errorf("countOneBits() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
