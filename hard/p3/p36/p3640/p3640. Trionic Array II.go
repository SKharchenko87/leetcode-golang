package p3640

import "math"

type SegmentSum struct {
	startPrefSum            int64
	endPrefSum              int64
	startNotNegativePrefSum int64
	last2Sum                int64
	increasing              bool
	lastNum                 int64
	first2SumFull           int64
}

// Селектор для хранения последних 3-х сегментов
type TrionicWindow struct {
	segments [3]*SegmentSum
	size     int
	end      int
}

func (tw *TrionicWindow) addSegment(seg *SegmentSum, start, end int) {
	defer func(e int) { tw.end = e }(end)

	if tw.size == 0 {
		tw.segments[0] = seg
		tw.size = 1
		return
	}

	if tw.end != start {
		tw.segments[0], tw.segments[1], tw.segments[2] = seg, nil, nil
		tw.size = 1
		return
	}

	if tw.size < 3 {
		tw.segments[tw.size] = seg
		tw.size++
		return
	}

	copy(tw.segments[0:], tw.segments[1:])
	tw.segments[2] = seg

}

func (tw TrionicWindow) check() bool {
	if tw.size < 3 {
		return false
	}
	s := tw.segments
	return s[0].increasing && !s[1].increasing && s[2].increasing
}

func (tw TrionicWindow) sum() int64 {
	s := tw.segments

	sum0 := max(
		s[0].endPrefSum-s[0].startPrefSum,
		s[0].endPrefSum-s[0].startNotNegativePrefSum,
		s[0].last2Sum,
	)

	sum1 := s[1].endPrefSum - s[1].startPrefSum

	sum2 := max(
		s[2].endPrefSum-s[2].startPrefSum+s[2].lastNum,
		s[2].first2SumFull,
	)

	return sum0 + sum1 + sum2
}

const (
	notHaveExtremum = iota
	minimumExtremum
	maximumExtremum
	flatExtremum
)

type extremum int

func hasExtremum(nums []int, i int) extremum {
	a, b, c := nums[i-1], nums[i], nums[i+1]
	if a == b {
		return flatExtremum
	}
	if a < b && b >= c {
		return maximumExtremum
	}
	if a > b && b <= c {
		return minimumExtremum
	}
	return notHaveExtremum
}

func maxSumTrionic(nums []int) int64 {
	res := int64(math.MinInt64)
	prevSum := int64(nums[0])
	hasNonNegative := false
	n := len(nums)
	startIndex := 0
	startPrefSum := int64(0)
	first2SumFull := int64(nums[0] + nums[1])
	tw := TrionicWindow{}
	var endPrefSum, last2Sum, startNotNegativePrefSum int64
	startNotNegativePrefSum = math.MaxInt64 / 2
	i := 1
	newSegmentSum := func(ext extremum) *SegmentSum {
		return &SegmentSum{
			startPrefSum:            startPrefSum,
			endPrefSum:              endPrefSum,
			startNotNegativePrefSum: startNotNegativePrefSum,
			last2Sum:                last2Sum,
			increasing:              ext == maximumExtremum,
			lastNum:                 int64(nums[i]),
			first2SumFull:           first2SumFull,
		}
	}
	var ext extremum
	reset := func() {
		if !hasNonNegative && nums[i] >= 0 {
			hasNonNegative = true
			startNotNegativePrefSum = prevSum
		}
		prevSum += int64(nums[i])
	}
	for ; i < n-1; i++ {
		switch ext = hasExtremum(nums, i); ext {
		case notHaveExtremum:
			reset()
			continue
		case minimumExtremum, maximumExtremum:

			endPrefSum = prevSum
			last2Sum = int64(nums[i-1])

			seg := newSegmentSum(ext)
			tw.addSegment(seg, startIndex, i)
			if tw.check() {
				res = max(res, tw.sum())
			}
		}
		hasNonNegative = false
		startNotNegativePrefSum = math.MaxInt64 / 2
		startPrefSum = prevSum
		first2SumFull = int64(nums[i] + nums[i+1])
		startIndex = i
		reset()
	}
	if nums[i-1] == nums[i] {
		return res
	}

	endPrefSum = prevSum
	last2Sum = int64(nums[i-1])

	ext = minimumExtremum
	if nums[i-1] < nums[i] {
		ext = maximumExtremum
	}
	seg := newSegmentSum(ext)

	tw.addSegment(seg, startIndex, i)
	if tw.check() {
		res = max(res, tw.sum())
	}

	return res
}

/**********************/
type Subarray struct {
	start, end       int
	startNotNegative int
	increasing       bool
}

func checkMax(sub []Subarray, prevSum []int) int64 {
	res := int64(math.MinInt64)
	if !(sub[0].end == sub[1].start && sub[1].end == sub[2].start && sub[0].increasing && !sub[1].increasing && sub[2].increasing) {
		return res
	}

	var sum00, sum01, sum02, sum1, sum20, sum21 int64
	var start, end int

	start = sub[0].start
	end = sub[0].end + 1
	sum00 = int64(prevSum[end] - prevSum[start])

	if sub[0].startNotNegative > 0 && sub[0].startNotNegative != sub[0].end {
		start = sub[0].startNotNegative
		end = sub[0].end + 1
		sum01 = int64(prevSum[end] - prevSum[start])
	} else {
		sum01 = math.MinInt64
	}

	start = sub[0].end - 1
	end = sub[0].end + 1
	sum02 = int64(prevSum[end] - prevSum[start])

	start = end
	end = sub[1].end + 1
	sum1 = int64(prevSum[end] - prevSum[start])

	start = end
	end = sub[2].start + 1 + 1
	sum20 = int64(prevSum[end] - prevSum[start])

	end = sub[2].end + 1
	sum21 = int64(prevSum[end] - prevSum[start])
	return max(sum00, sum01, sum02) + sum1 + max(sum20, sum21)
}

func maxSumTrionic0(nums []int) int64 {
	n := len(nums)
	prevSum := make([]int, n+1)
	prevSum[1] = nums[0]
	for i := 1; i < n; i++ {
		prevSum[i+1] = prevSum[i] + nums[i]
	}

	start := 0
	startNotNegative := -1
	subArr := make([]Subarray, 0, 4)
	isUp := true
	i := 1
	res := int64(math.MinInt64)
	for ; i < n; i++ {
		if nums[i] > nums[i-1] {
			// Проверяем на экстремум
			if i-1-start > 0 {
				isUp = nums[i-1] > nums[i-2]
				if !isUp {
					sa := Subarray{
						start:            start,
						end:              i - 1,
						startNotNegative: startNotNegative,
						increasing:       false,
					}
					subArr = append(subArr, sa)
					if len(subArr) == 4 {
						subArr = subArr[1:]
					}
					if len(subArr) == 3 {
						res = max(res, checkMax(subArr, prevSum))
					}
					start = i - 1
					startNotNegative = -1
				}
			}
			if startNotNegative == -1 && nums[i] >= 0 {
				startNotNegative = i
			}
		} else if nums[i] == nums[i-1] {
			if i-1-start > 0 {
				sa := Subarray{
					start:            start,
					end:              i - 1,
					startNotNegative: startNotNegative,
					increasing:       nums[start] < nums[i-1],
				}
				subArr = append(subArr, sa)
				if len(subArr) == 4 {
					subArr = subArr[1:]
				}
				if len(subArr) == 3 {
					res = max(res, checkMax(subArr, prevSum))
				}
			}
			startNotNegative = -1
			start = i
		} else if nums[i] < nums[i-1] {
			// Проверяем на экстремум
			if i-1-start > 0 {
				isDown := nums[i-1] < nums[i-2]
				if !isDown {
					sa := Subarray{
						start:            start,
						end:              i - 1,
						startNotNegative: startNotNegative,
						increasing:       true,
					}
					subArr = append(subArr, sa)
					if len(subArr) == 4 {
						subArr = subArr[1:]
					}
					if len(subArr) == 3 {
						res = max(res, checkMax(subArr, prevSum))
					}

					start = i - 1
					startNotNegative = -1
				}
			}
		}
	}
	if i-1-start > 0 && nums[i-1] != nums[i-2] {
		sa := Subarray{
			start:            start,
			end:              i - 1,
			startNotNegative: startNotNegative,
			increasing:       nums[i-1] > nums[i-2],
		}
		subArr = append(subArr, sa)
		if len(subArr) == 4 {
			subArr = subArr[1:]
		}
		if len(subArr) == 3 {
			res = max(res, checkMax(subArr, prevSum))
		}

	}

	return res
}
