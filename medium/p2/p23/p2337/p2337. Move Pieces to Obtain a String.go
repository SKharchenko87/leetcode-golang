package p2337

func canChange(start string, target string) bool {
	ls, lt := len(target), len(start)
	index0, index1 := 0, 0
	for {
		for ; index0 < ls && start[index0] == '_'; index0++ {
		}
		for ; index1 < lt && target[index1] == '_'; index1++ {
		}
		if index0 == ls || index1 == lt {
			break
		} else if start[index0] != target[index1] {
			return false
		} else if start[index0] == 'L' && index0 < index1 {
			return false
		} else if start[index0] == 'R' && index0 > index1 {
			return false
		}
		index0++
		index1++
	}
	return index0 == index1
}
