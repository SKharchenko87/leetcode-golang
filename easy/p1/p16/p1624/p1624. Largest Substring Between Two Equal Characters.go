package p1624

import "fmt"

func maxLengthBetweenEqualCharacters(s string) int {
	m := map[int32][]int{'a': []int{}, 'b': []int{}, 'c': []int{}, 'd': []int{}, 'e': []int{}, 'f': []int{}, 'g': []int{}, 'h': []int{}, 'i': []int{}, 'j': []int{}, 'k': []int{}, 'l': []int{}, 'm': []int{}, 'n': []int{}, 'o': []int{}, 'p': []int{}, 'q': []int{}, 'r': []int{}, 's': []int{}, 't': []int{}, 'u': []int{}, 'v': []int{}, 'w': []int{}, 'x': []int{}, 'y': []int{}, 'z': []int{}}
	maxRes := -1
	for i, c := range s {
		m[c] = append(m[c], i)
		l := len(m[c])
		if l > 1 && maxRes < i-m[c][0]-1 {
			maxRes = i - m[c][0] - 1
		}
	}
	fmt.Println(m)
	return maxRes
}
