package p0725

import . "leetcode/stucture"

func splitListToParts(head *ListNode, k int) []*ListNode {
	listSize := 0
	res := make([]*ListNode, k)
	res[0] = head
	flgNull := false
	var partSize, indexOfPart, extendedPartCount int
	for head != nil {
		defer func(i int, head *ListNode) {
			if flgNull {
				head.Next = nil
				flgNull = false
			}
			indexOfPart--
			if indexOfPart == 0 && k > 1 {
				k--
				res[k] = head
				indexOfPart = partSize
				if k <= extendedPartCount {
					indexOfPart = partSize + 1
				}
				flgNull = true
			}
		}(listSize, head)
		head = head.Next
		listSize++
	}
	extendedPartCount = listSize % k
	partSize = listSize / k
	if k > listSize {
		k = listSize
		extendedPartCount = 0
	}
	if partSize == 0 {
		partSize++
	}
	indexOfPart = partSize
	return res
}

func run(head []int, k int) [][]int {
	tmp := splitListToParts(ArrToList(head), k)
	res := [][]int{}
	for _, node := range tmp {
		res = append(res, ListToArr(node))
	}
	return res
}

// 0,1,2,3, 4,5,6, 7,8,9
