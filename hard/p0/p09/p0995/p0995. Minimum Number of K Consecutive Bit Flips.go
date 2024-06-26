package p0995

import "fmt"

func minKBitFlips(nums []int, k int) int {
	l := len(nums)
	cnt := 0
	currentFlips := 0
	for i := 0; i < l; i++ {
		if i >= k && nums[i-k] == 2 {
			currentFlips--
		}
		if nums[i] == currentFlips&1 {
			currentFlips++
			cnt++
			if i+k > l {
				return -1
			}
			nums[i] = 2
		}
	}
	return cnt
}

func minKBitFlips4(nums []int, k int) int {
	l := len(nums)
	cnt := 0
	prevSum := 0
	index := 0
	lastFlip := -1
	for i := 0; i < l; i++ {
		if lastFlip >= 0 && nums[index] == i-k && prevSum > 0 {
			prevSum--
			index++
		}
		if nums[i] == prevSum&1 {
			prevSum++
			cnt++
			if i+k > l {
				return -1
			}
			lastFlip++
			nums[lastFlip] = i
		}
	}
	return cnt
}

func minKBitFlips3(nums []int, k int) int {
	l := len(nums)
	cnt := 0
	prevSum := 0
	index := 0
	lastFlip := -1
	for i := 0; i < l; i++ {
		if lastFlip >= 0 && nums[index] == i-k && prevSum > 0 {
			prevSum--
			index++
		}
		if nums[i] == 0 && prevSum&1 == 0 || nums[i] == 1 && prevSum&1 == 1 {
			prevSum++
			cnt++
			if i+k > l {
				return -1
			}
			lastFlip++
			nums[lastFlip] = i
		}
	}
	return cnt
}

func minKBitFlips2(nums []int, k int) int {
	l := len(nums)
	cnt := 0
	prevSum := 0
	indexArr := make([]int, l-k+1)
	index := 0
	i2 := -1
	for i := 0; i < l; i++ {
		if i2 >= 0 && indexArr[index] == i-k {
			prevSum--
			index++
		}
		if nums[i] == 0 && prevSum&1 == 0 || nums[i] == 1 && prevSum&1 == 1 {
			prevSum++
			cnt++
			if i+k > l {
				return -1
			}
			i2++
			indexArr[i2] = i
		}
	}
	return cnt
}

func minKBitFlips1(nums []int, k int) int {
	l := len(nums)
	cnt := 0
	prevSum := 0
	indexArr := []int{}
	for i := 0; i < l; i++ {
		if len(indexArr) > 0 && indexArr[0] == i-k {
			prevSum--
			indexArr = indexArr[1:]
		}
		if nums[i] == 0 && prevSum&1 == 0 || nums[i] == 1 && prevSum&1 == 1 {
			prevSum++
			cnt++
			if i+k > l {
				return -1
			}
			indexArr = append(indexArr, i)
		}
		print(prevSum, " ")
	}
	fmt.Println(indexArr)
	return cnt
}

func minKBitFlips0(nums []int, k int) int {
	l := len(nums)
	cnt := 0
	for i := 0; i < l-k+1; i++ {
		if nums[i] == 0 {
			for j := i; j-i < k; j++ {
				nums[j] = 1 - nums[j]
			}
			cnt += 1
		}
	}
	for i := l - k; i < l; i++ {
		if nums[i] == 0 {
			return -1
		}
	}
	return cnt
}
