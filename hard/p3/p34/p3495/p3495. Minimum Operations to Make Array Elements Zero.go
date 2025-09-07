package p3495

import (
	"slices"
)

type four struct {
	i      int64
	sum    int64
	degree int64
}

var data []four = []four{{0, 0, 0}, {1, 1, 1}, {4, 5, 2}, {16, 30, 3}, {64, 175, 4}, {256, 944, 5}, {1024, 4785, 6}, {4096, 23218, 7}, {16384, 109235, 8}, {65536, 502452, 9}, {262144, 2271925, 10}, {1048576, 10136246, 11}, {4194304, 44739255, 12}, {16777216, 195734200, 13}, {67108864, 850045625, 14}, {268435456, 3668617914, 15}}

func minOperations(queries [][]int) int64 {
	res := int64(0)
	for _, q := range queries {
		tmp := operation(int64(q[1])) - operation(int64(q[0]-1))
		d := tmp % 2
		tmp /= 2
		res += tmp + d
	}
	return res
}

func operation(x int64) int64 {
	index, ok := slices.BinarySearchFunc(data, four{x, 0, 0}, func(a, b four) int {
		return int(a.i - b.i)
	})
	if ok {
		return data[index].sum
	}
	return data[index-1].sum + (x-data[index-1].i)*data[index-1].degree
}

//
//func ppp() int {
//	c := 257
//	res := make([]int, c)
//	for i := 1; i < c; i++ {
//		res[i] = res[i-1] + operation(i)
//	}
//	//fmt.Println(res)
//	for i := 0; i < len(res); i++ {
//		if i == 0 || i == 4 || i == 16 || i == 64 || i == 256 {
//			fmt.Println(i, res[i])
//		}
//	}
//	prev := int64(0)
//	prevPow := int64(0)
//	for i := int64(1); i < 1e10; i *= 4 {
//		prev = prev + (i-i/4)*prevPow
//		prevPow++
//		fmt.Println(i, prev+prevPow, prevPow)
//	}
//	return 0
//}
