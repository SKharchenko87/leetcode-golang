package p1345

import "testing"

func Test_minJumps(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arr: []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}}, 3},
		{"Example 2", args{arr: []int{7}}, 0},
		{"Example 3", args{arr: []int{7, 6, 9, 6, 9, 6, 9, 7}}, 1},
		{"TestCase 25", args{arr: []int{51, 64, -15, 58, 98, 31, 48, 72, 78, -63, 92, -5, 64, -64, 51, -48, 64, 48, -76, -86, -5, -64, -86, -47, 92, -41, 58, 72, 31, 78, -15, -76, 72, -5, -97, 98, 78, -97, -41, -47, -86, -97, 78, -97, 58, -41, 72, -41, 72, -25, -76, 51, -86, -65, 78, -63, 72, -15, 48, -15, -63, -65, 31, -41, 95, 51, -47, 51, -41, -76, 58, -81, -41, 88, 58, -81, 88, 88, -47, -48, 72, -25, -86, -41, -86, -64, -15, -63}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minJumps(tt.args.arr); got != tt.want {
				t.Errorf("minJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
