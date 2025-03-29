package p2818

import "sort"

const mx = 100_000
const mod = 1_000_000_007
const inf = int(1e18)

var primes [mx + 1]int

func init() {
	primes[0] = inf
	for i := 2; i <= mx; i++ {
		if primes[i] == 0 {
			for j := i; j <= mx; j += i {
				primes[j]++
			}
		}
	}
}

func maximumScoreX(nums []int, k int) int {
	nums = append(nums, 0)
	stack, count := []int{-1}, make(map[int]int)
	for i, v := range nums {
		for len(stack) > 1 && primes[v] > primes[nums[stack[len(stack)-1]]] {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			count[nums[p]] += (p - stack[len(stack)-1]) * (i - p)
		}
		stack = append(stack, i)
	}
	keys := []int{}
	for v := range count {
		keys = append(keys, v)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	res := 1
	for i := 0; k > 0; i++ {
		t := min(count[keys[i]], k)
		k -= t
		res = res * pow(keys[i], t) % mod
	}
	return res
}
