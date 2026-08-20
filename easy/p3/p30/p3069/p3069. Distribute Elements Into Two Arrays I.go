package p3069

func resultArray(nums []int) []int {
	n := len(nums)
	arr1, arr2 := make([]int, 1, n), make([]int, 1, n/2)
	arr1[0] = nums[0]
	arr2[0] = nums[1]

	for i := 2; i < len(nums); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	return append(arr1, arr2...)
}
