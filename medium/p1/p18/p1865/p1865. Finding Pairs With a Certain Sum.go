package p1865

type FindSumPairs struct {
	nums1    []int
	nums2    []int
	nums2Map map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	nums2Map := map[int]int{}
	for _, num2 := range nums2 {
		nums2Map[num2]++
	}

	return FindSumPairs{nums1, nums2, nums2Map}
}

func (this *FindSumPairs) Add(index int, val int) {

	oldVal := this.nums2[index]
	newVal := oldVal + val

	this.nums2Map[oldVal]--
	this.nums2Map[newVal]++
	this.nums2[index] = newVal
}

func (this *FindSumPairs) Count(tot int) int {
	res := 0
	for _, num1 := range this.nums1 {
		res += this.nums2Map[tot-num1]
	}
	return res
}

func run(commands []string, params []any) []any {
	result := make([]any, len(commands))
	var obj FindSumPairs
	for i, command := range commands {
		switch command {
		case "FindSumPairs":
			param := params[i].([][]int)
			obj = Constructor(param[0], param[1])
			result[i] = nil
		case "count":
			param := params[i].([]int)
			result[i] = obj.Count(param[0])
		case "add":
			param := params[i].([]int)
			obj.Add(param[0], param[1])
			result[i] = nil
		}
	}
	return result
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
