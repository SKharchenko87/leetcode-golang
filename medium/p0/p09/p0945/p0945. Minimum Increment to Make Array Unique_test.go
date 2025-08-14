package p0945

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_minIncrementForUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	println(filepath.Abs("./60.data"))
	case60 := arrayFromFileRead("./60.data")
	case61 := arrayFromFileRead("./61.data")
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 2, 2}}, 1},
		{"Example 2", args{[]int{3, 2, 1, 2, 1, 7}}, 6},
		{"Example 35", args{[]int{0, 0}}, 1},
		{"Example 60", args{case60}, 96896},
		{"Example 61", args{case61}, 191376},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIncrementForUnique(tt.args.nums); got != tt.want {
				t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func arrayFromFileRead(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var perline int
	var nums []int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}
	// print out the nums array content
	return nums
}
