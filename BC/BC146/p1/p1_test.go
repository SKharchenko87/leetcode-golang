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

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 1, 4, 1}}, 1},
		{"Example 2", args{nums: []int{1, 1, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
