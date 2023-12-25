package main

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
	if s.count != 0 {
		if len(s.arr) <= s.count {
			s.arr = append(s.arr, e)
		} else {
			s.arr[s.count-1] = e
		}
		s.count++
	} else {
		if len(s.arr) > 0 {
			s.arr[0] = e
		} else {
			s.arr = append(s.arr, e)
		}
		s.count = 1

	}
}

func (s *stack[T]) peek() T {
	return s.arr[s.count-1]
}

func (s *stack[T]) getCount() int {
	return s.count
}

func (s *stack[T]) pop() T {
	x := s.arr[s.count-1]
	s.count--
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
	stackBracket := stack[int32]{}
	stackInt := stack[int]{}

	bufStr := []int32{}
	bufInt := []int32{}

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			bufStr = append(bufStr, v)
		}
		if v >= '0' && v <= '9' {
			if len(bufStr) > 0 {
				stackString.push(string(bufStr))
				bufStr = []int32{}
			}
			bufInt = append(bufInt, v)
		}
		if v == '[' {
			stackBracket.push(v)
			stackInt.push(arrToInt(bufInt))
			bufInt = []int32{}
		}
		if v == ']' {
			var str string
			n := stackInt.pop()
			if len(bufStr) > 0 {
				str = strings.Repeat(string(bufStr), n)
			} else {
				str = strings.Repeat(stackString.pop(), n)
			}
			stackBracket.pop()
			if stackBracket.getCount() > 0 {
				str = stackString.pop() + str
			}
			stackString.push(str)
			bufStr = []int32{}
		}
	}
	sss := ""
	for stackString.getCount() != 0 {
		sss = stackString.pop() + sss
	}
	if len(bufStr) > 0 {
		sss = sss + string(bufStr)
	}
	return sss
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
