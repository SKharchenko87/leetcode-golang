package p2333

import "sort"

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	// Размер n+1 — отличный трюк! В конце будет 0,
	// к которому мы будем стремиться, если k очень большое.
	nums := make([]int, n+1)

	for i := 0; i < n; i++ {
		diff := nums1[i] - nums2[i]
		if diff < 0 {
			diff = -diff
		}
		nums[i] = diff
	}

	// Сортируем по убыванию
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	k := int64(k1 + k2)

	for i := 1; i <= n; i++ {
		if nums[i-1] > nums[i] {
			count := int64(i)
			gap := int64(nums[i-1] - nums[i])
			cost := count * gap

			if k >= cost {
				// Шагов хватает, чтобы "срезать" все верхние элементы до уровня nums[i]
				k -= cost
			} else {
				// Шагов не хватает. Распределяем k поровну между верхними i элементами
				drop := k / count // На столько гарантированно уменьшится каждый из i элементов
				rem := k % count  // Столько элементов уменьшатся еще на 1

				ans := int64(0)
				val1 := int64(nums[i-1]) - drop - 1
				val2 := int64(nums[i-1]) - drop

				// Добавляем квадраты тех, что уменьшились на (drop + 1)
				ans += rem * val1 * val1
				// Добавляем квадраты тех, что уменьшились только на drop
				ans += (count - rem) * val2 * val2

				// Добавляем квадраты всех остальных нетронутых элементов
				for j := i; j < n; j++ {
					ans += int64(nums[j]) * int64(nums[j])
				}

				return ans
			}
		}
	}

	// Если цикл прошел до конца и мы вышли из него,
	// значит мы смогли свести все разницы к нулю
	return 0
}

/*TLE*/
func minSumSquareDiff0(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	nums := make([]int, n+1)
	for i := 0; i < n; i++ {
		nums[i] = abs(nums1[i] - nums2[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	j := 0
	for i := 0; i < k1+k2 && nums[j] > 0; i++ {
		nums[j]--
		if nums[j] >= nums[j+1] {
			j = 0
			continue
		}
		j++
	}
	res := int64(0)
	for i := 0; i < n; i++ {
		res += int64(nums[i] * nums[i])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
