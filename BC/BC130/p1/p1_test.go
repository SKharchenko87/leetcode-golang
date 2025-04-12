package p1

import "testing"

//func Test_abs(t *testing.T) {
//	type args struct {
//		x int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{"Case 1", args{-1}, 1},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := abs(tt.args.x); got != tt.want {
//				t.Errorf("abs() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_abs(t *testing.T) {
	type args[T Numeric] struct {
		x T
	}
	type testCase[T Numeric] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
		{name: "Test 1", args: args[int]{x: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_satisfiesConditions(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{grid: [][]int{{1, 0, 2}, {1, 0, 2}}}, true},
		{"Example 2", args{grid: [][]int{{1, 1, 1}, {0, 0, 0}}}, false},
		{"Example 3", args{grid: [][]int{{1}, {2}, {3}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := satisfiesConditions(tt.args.grid); got != tt.want {
				t.Errorf("satisfiesConditions() = %v, want %v", got, tt.want)
			}
		})
	}
}
