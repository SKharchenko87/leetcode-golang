package p0394

import (
	"fmt"
	"strings"
)

func main() {
	//s := "3[a]2[bc]"
	//s := "2[abc]3[cd]ef"
	//s := "qqq2[abc]3[cd]ef"
	//s := "100[leetcode]"
	s := "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
	fmt.Println(decodeString(s))
}

type stack[T any] struct {
	arr   []T
	count int
}

func (s *stack[T]) push(e T) {
	s.arr = append(s.arr, e)
	s.count = len(s.arr)
}

func (s *stack[T]) peek() T {
	return s.arr[s.count-1]
}

func (s *stack[T]) getCount() int {
	return s.count
}

func (s *stack[T]) pop() T {
	var nil T // Type must be a pointer, channel, func, interface, map, or slice type
	if s.count == 0 {
		return nil
	}
	x := s.arr[s.count-1]
	s.arr = s.arr[:s.count-1]
	s.count = len(s.arr)
	return x
}

func arrToInt(arr []int32) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res = res*10 + int(arr[i]-'0')
	}
	return res
}

func decodeString(s string) string {
	stackString := stack[string]{}
	stackInt := stack[int]{}

	bufStr := []int32{}
	bufInt := []int32{}
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			bufStr = append(bufStr, v)
		}
		if v >= '0' && v <= '9' {
			bufInt = append(bufInt, v)
		}
		if v == '[' {
			stackInt.push(arrToInt(bufInt))
			bufInt = []int32{}
			stackString.push(string(bufStr))
			bufStr = []int32{}
		}
		if v == ']' {
			x := stackInt.pop()
			var tmp, repeatStr string
			if len(bufStr) == 0 {
				tmp = stackString.pop()
			} else {
				tmp = string(bufStr)
			}
			repeatStr = stackString.pop() + strings.Repeat(tmp, x)
			bufStr = []int32(repeatStr)
		}
	}

	res := ""
	for stackString.getCount() != 0 {
		res = stackString.pop() + res
	}
	if len(bufStr) > 0 {
		res = res + string(bufStr)
	}
	return res
}

func repeatArr[T any](arr []T, n int) []T {
	arrLength := len(arr)
	resLength := arrLength * n
	res := make([]T, resLength)
	for i := 0; i < resLength; i++ {
		res[i] = arr[i%arrLength]
	}
	return res
}
