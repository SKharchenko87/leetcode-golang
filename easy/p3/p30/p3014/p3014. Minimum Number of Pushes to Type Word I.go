package p3014

import (
	"slices"
	"sort"
)

func minimumPushes(word string) int {
	alphabet := make([]int, 26)
	for i := 0; i < len(word); i++ {
		alphabet[int(word[i])-'a']++
	}
	sort.Ints(alphabet)
	res := 0
	for i := 25; i >= 0; i-- {
		res += alphabet[i] * ((25-i)/8 + 1)
	}
	return res
}

func minimumPushes3(word string) int {
	alphabet := make([]int, 26)
	for i := 0; i < len(word); i++ {
		alphabet[int(word[i])-'a']++
	}
	sort.Ints(alphabet)
	res := 0
	for i, x := range slices.Backward(alphabet) {
		res += x * ((25-i)/8 + 1)
	}
	return res
}

func minimumPushes2(word string) int {
	alphabet := make([]int, 26)
	for i := 0; i < len(word); i++ {
		alphabet[int(word[i])-'a']++
	}
	sort.Ints(alphabet)
	res := 0
	for i := 25; i >= 0; i-- {
		switch {
		case i < 2:
			res += alphabet[i] * 4
		case i < 10:
			res += alphabet[i] * 3
		case i < 18:
			res += alphabet[i] * 2
		case i < 26:
			res += alphabet[i]
		}
	}
	return res
}

func minimumPushes1(word string) int {
	alphabet := make([]int, 26)
	for i := 0; i < len(word); i++ {
		alphabet[int(word[i])-'a']++
	}
	sort.Ints(alphabet)
	res := 0
	i := 0
	for ; i < 8; i++ {
		res += alphabet[25-i]
	}
	for ; i < 16; i++ {
		res += alphabet[25-i] * 2
	}
	for ; i < 24; i++ {
		res += alphabet[25-i] * 3
	}
	for ; i < 26; i++ {
		res += alphabet[25-i] * 4
	}
	return res
}

func minimumPushes0(word string) int {
	arr := make([][2]int, 26)
	m := map[byte]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
	l := len(word)
	for i := 0; i < l; i++ {
		m[word[i]]++
		arr[word[i]-'a'] = [2]int{int(word[i]), m[word[i]]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	count := 0
	for i, v := range arr {
		k := i/8 + 1
		count = count + v[1]*k
	}
	return count
}
