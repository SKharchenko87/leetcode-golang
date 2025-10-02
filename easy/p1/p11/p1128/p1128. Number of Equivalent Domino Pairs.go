package p1128

func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[int]int)
	res := 0
	for _, dominoe := range dominoes {
		var val int
		if dominoe[0] > dominoe[1] {
			val = dominoe[1]*10 + dominoe[0]
		} else {
			val = dominoe[0]*10 + dominoe[1]
		}
		res += m[val]
		m[val]++
	}
	return res
}

func numEquivDominoPairs2(dominoes [][]int) int {
	m := make(map[[2]int]int)
	res := 0
	for _, dominoe := range dominoes {
		val := [2]int{min(dominoe[1], dominoe[0]), max(dominoe[1], dominoe[0])}
		res += m[val]
		m[val]++
	}
	return res
}

func numEquivDominoPairs1(dominoes [][]int) int {
	l := len(dominoes)
	m := make(map[[2]int]int, l)
	res := 0
	for _, dominoe := range dominoes {
		var val [2]int
		if dominoe[0] > dominoe[1] {
			val = [2]int{dominoe[1], dominoe[0]}
		} else {
			val = [2]int{dominoe[0], dominoe[1]}
		}
		res += m[val]
		m[val]++
	}
	return res
}

func numEquivDominoPairs0(dominoes [][]int) int {
	l := len(dominoes)
	m := make(map[int]int, l)
	res := 0
	for _, dominoe := range dominoes {
		var val int
		if dominoe[0] > dominoe[1] {
			val = dominoe[1]*10 + dominoe[0]
		} else {
			val = dominoe[0]*10 + dominoe[1]
		}
		res += m[val]
		m[val]++
	}
	return res
}
