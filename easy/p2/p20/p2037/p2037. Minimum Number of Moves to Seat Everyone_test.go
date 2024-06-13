package p2037

import "testing"

func Test_minMovesToSeat(t *testing.T) {
	type args struct {
		seats    []int
		students []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"Example 1", args{seats: []int{3, 1, 5}, students: []int{2, 7, 4}}, 4},
		{"Example 2", args{seats: []int{4, 1, 5, 9}, students: []int{1, 3, 2, 6}}, 7},
		{"Example 3", args{seats: []int{2, 2, 6, 6}, students: []int{1, 3, 2, 6}}, 4},
		{"Example 53", args{seats: []int{12, 14, 19, 19, 12}, students: []int{19, 2, 17, 20, 7}}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := minMovesToSeat(tt.args.seats, tt.args.students); gotRes != tt.wantRes {
				t.Errorf("minMovesToSeat() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
