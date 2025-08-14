package p2163

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func Test_minimumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{3, 1, 2}}, -1},
		{"Example 2", args{nums: []int{7, 9, 5, 8, 1, 3}}, 1},
		{"TestCase 21", args{nums: []int{16, 46, 43, 41, 42, 14, 36, 49, 50, 28, 38, 25, 17, 5, 18, 11, 14, 21, 23, 39, 23}}, -14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_z(t *testing.T) {
	arr := []int{
		16, 46, 43, 41, 42, 14, 36,
		//49,
		//50,
		//28,
		//38,
		//25,
		//17,
		//5,
	}
	sort.Ints(arr)
	//slices.Reverse(arr)
	sum := 0
	for i := 0; i < 7; i++ {
		sum += arr[i]
	}

	arr0 := []int{
		49,
		50,
		28,
		38,
		25,
		17,
		5,
		18, 11, 14, 21, 23, 39, 23,
	}
	sort.Ints(arr0)
	slices.Reverse(arr0)
	sum0 := 0
	for i := 0; i < 7; i++ {
		sum0 += arr0[i]
	}

	fmt.Println(sum - sum0)

	fmt.Println(arr)
	fmt.Println(sum)

	fmt.Println(arr0)
	fmt.Println(sum0)

}
