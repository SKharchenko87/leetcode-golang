package p0621

import "sort"

func leastInterval(tasks []byte, n int) int {
	maxFreq := 0
	dict := make(map[byte]int)

	for i := 0; i < len(tasks); i++ {
		dict[tasks[i]]++
		if dict[tasks[i]] > maxFreq {
			maxFreq = dict[tasks[i]]
		}
	}
	res := (maxFreq - 1) * (n + 1)
	for _, value := range dict {
		if value == maxFreq {
			res++
		}
	}
	return max(res, len(tasks))
}

func leastInterval0(tasks []byte, n int) int {
	arr := make([]int, 28)
	for _, task := range tasks {
		arr[task-'A']++
	}
	sort.Ints(arr)
	df := 0
	res := 0
	for i := 1; i < 28; i++ {
		if arr[i] != arr[i-1] {
			if arr[i] == arr[27] {
				res += (arr[i] - arr[i-1] - 1) * (n + 1)
				res += 28 - i
				break
			}
			if 28-i > n+1 {
				df = (28 - i)
			} else {
				df = (n + 1)
			}
			res += (arr[i] - arr[i-1]) * df
		}
	}
	return res
}
