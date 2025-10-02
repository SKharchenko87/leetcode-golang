package p3016

import "sort"

func minimumPushes(word string) int {
	arr := make([]int, 26)
	for _, ch := range word {
		arr[ch-'a']++
	}
	sort.Ints(arr)
	var result int
	coefficient := 1
	i := 25
	for ; i > 17; i-- {
		result += arr[i]
	}
	coefficient++
	for ; i > 9; i-- {
		result += arr[i] * 2
	}
	coefficient++
	for ; i > 1; i-- {
		result += arr[i] * 3
	}
	coefficient++
	for ; i >= 0; i-- {
		result += arr[i] * 4
	}
	return result
}

func minimumPushes1(word string) int {
	arr := make([]int, 26)
	for _, ch := range word {
		arr[ch-'a']++
	}
	sort.Ints(arr)
	var result int
	coefficient := 1
	for i := 25; i > 0 && arr[i] > 0; i-- {
		if i == 17 || i == 9 || i == 1 {
			coefficient++
		}
		result += arr[i] * coefficient
	}
	return result
}

func minimumPushes0(word string) int {
	arr := make([]int, 26)
	for i := 0; i < len(word); i++ {
		arr[word[i]-'a']++
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	count := 0
	for i, v := range arr {
		count += v * (i/8 + 1)
	}
	return count
}
