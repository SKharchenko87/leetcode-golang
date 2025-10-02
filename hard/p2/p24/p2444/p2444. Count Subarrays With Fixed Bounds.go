package p2444

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64 = 0
	minPos, maxPos, leftBound := -1, -1, -1

	for i := 0; i < len(nums); i++ {
		// Если текущий элемент вне диапазона, обновляем leftBound
		if nums[i] < minK || nums[i] > maxK {
			leftBound = i
		}

		// Обновляем позиции minK и maxK
		if nums[i] == minK {
			minPos = i
		}
		if nums[i] == maxK {
			maxPos = i
		}

		// Вычисляем количество валидных подмассивов, заканчивающихся на i
		validSubarrays := int64(min(minPos, maxPos)) - int64(leftBound)
		if validSubarrays > 0 {
			res += validSubarrays
		}
	}

	return res
}

// 1, 3, 5, 2, 7, 5
// m  0  m  0  -  m
