package p3495

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{queries: [][]int{{1, 2}, {2, 4}}}, 3},
		{"Example 2", args{queries: [][]int{{2, 6}}}, 4},
		{"TestCase 11", args{queries: [][]int{{1, 8}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.queries); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func Test_ppp(t *testing.T) {
//	tests := []struct {
//		name string
//		want int
//	}{
//		{"Example 1", 1},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ppp(); got != tt.want {
//				t.Errorf("ppp() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
