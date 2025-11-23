package p2654

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minOperations(nums []int) int {
	res := len(nums)
	for _, num := range nums {
		if num == 1 {
			res--
		}
	}
	l := len(nums)
	if res < l {
		return res
	}

	for i := range nums {
		for j := 0; j < l-i-1; j++ {
			nums[j] = gcd(nums[j], nums[j+1])
			if nums[j] == 1 {
				return l + i
			}
		}
	}

	return -1
}

func minOperations0(nums []int) int {
	l := len(nums)
	ones := 0
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			ones++
		}
	}
	if ones != 0 {
		return l - ones
	}
	for i := 1; i < l; i++ {
		for j := 0; j < l-i; j++ {
			nums[j] = gcd(nums[j], nums[j+1])
			if nums[j] == 1 {
				return l + i - 1
			}
		}
	}
	return -1
}

/*
 2 6 9 6 2
  2 3 3 2
   1

gcd(2,6)=2 => 2->2 уже
gcd(6,9)=3 => 6->3 [2 3 9 6 2] (1)
gcd(2,3)=1 => 2->1 [1 3 9 6 2] (2)
gcd(1,9)=1 => 9->1 [1 1 1 6 2] (3)
gcd(1,6)=1 => 6->1 [1 1 1 1 2] (4)
gcd(1,2)=1 => 2->1 [1 1 1 1 1] (5)

2  42 10 12 30  15 42
  2  2  2  6  15  3
    2  2  2  3   3
      2  2  1  3


2  42 10 12 30  15 42
2  42 10  2 30  15 42
2  42 10  2  2  15 42
2  42 10  2  2   1 42


10 12 30  15
 2 12 30  15
 2 12 15  15
 2  3 15  15
 2  3  3  15
 1


10 12 30  15
10 12 15  15
10  3 15  15
 1  3 15  15


10 12 30  15
 2 12 30  15
 2  2 30  15
 2  2  2  15
 2  2  2   1



4, 2, 6, 3
 2  2  3
      1
*/
