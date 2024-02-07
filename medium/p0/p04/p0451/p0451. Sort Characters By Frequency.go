package p0451

import (
	"sort"
)

type Char struct {
	ch  byte
	cnt int
}

func frequencySort(s string) string {
	arr := []Char{{'a', 0}, {'b', 0}, {'c', 0}, {'d', 0}, {'e', 0}, {'f', 0}, {'g', 0}, {'h', 0}, {'i', 0}, {'j', 0}, {'k', 0}, {'l', 0}, {'m', 0}, {'n', 0}, {'o', 0}, {'p', 0}, {'q', 0}, {'r', 0}, {'s', 0}, {'t', 0}, {'u', 0}, {'v', 0}, {'w', 0}, {'x', 0}, {'y', 0}, {'z', 0}, {'A', 0}, {'B', 0}, {'C', 0}, {'D', 0}, {'E', 0}, {'F', 0}, {'G', 0}, {'H', 0}, {'I', 0}, {'J', 0}, {'K', 0}, {'L', 0}, {'M', 0}, {'N', 0}, {'O', 0}, {'P', 0}, {'Q', 0}, {'R', 0}, {'S', 0}, {'T', 0}, {'U', 0}, {'V', 0}, {'W', 0}, {'X', 0}, {'Y', 0}, {'Z', 0}, {'0', 0}, {'1', 0}, {'2', 0}, {'3', 0}, {'4', 0}, {'5', 0}, {'6', 0}, {'7', 0}, {'8', 0}, {'9', 0}}
	for _, c := range s {
		if c <= '9' {
			arr[52+c-'0'].cnt++
		} else if c <= 'Z' {
			arr[26+c-'A'].cnt++
		} else if c <= 'z' {
			arr[0+c-'a'].cnt++
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].cnt > arr[j].cnt {
			return true
		} else if arr[i].cnt == arr[j].cnt {
			return arr[i].ch < arr[j].ch
		}
		return false
	})
	resByte := make([]byte, len(s))
	k := 0
	for i := 0; i < 62 && arr[i].cnt != 0; i++ {
		for j := 0; j < arr[i].cnt; j++ {
			resByte[k] = arr[i].ch
			k++
		}
	}
	return string(resByte)
}
