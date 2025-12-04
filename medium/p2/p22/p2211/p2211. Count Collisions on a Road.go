package p2211

import "strings"

func countCollisions(directions string) int {
	var i, j int
	j = len(directions)
	for ; i < j && directions[i] == 'L'; i++ {
	}
	j--
	for ; 0 <= j && directions[j] == 'R'; j-- {
	}
	if i == -1 || j == -1 {
		return 0
	}
	return j - i - strings.Count(directions[i:j+1], "S") + 1
}

func countCollisions5(directions string) int {
	var i, j int
	j = len(directions) - 1
	for ; i <= j && directions[i] == 'L'; i++ {
	}
	for ; 0 <= j && directions[j] == 'R'; j-- {
	}
	if i == -1 || j == -1 {
		return 0
	}
	return j - i - strings.Count(directions[i:j+1], "S") + 1
}

func countCollisions4(directions string) int {
	var i, j int
	j = len(directions)
	for ; i < j && directions[i] == 'L'; i++ {
	}
	j--
	for ; 0 <= j && directions[j] == 'R'; j-- {
	}
	if i == -1 || j == -1 {
		return 0
	}
	countS := 0
	l := j - i + 1
	for ; i <= j; i++ {
		if directions[i] == 'S' {
			countS++
		}
	}
	return l - countS
}

func countCollisions3(directions string) int {
	i := strings.IndexAny(directions, "RS")
	j := strings.LastIndexAny(directions, "LS")
	if i == -1 || j == -1 {
		return 0
	}
	return j - i - strings.Count(directions[i:j+1], "S") + 1
}

func countCollisions2(directions string) int {
	i := strings.IndexAny(directions, "RS")
	j := strings.LastIndexAny(directions, "LS")
	if i == -1 || j == -1 {
		return 0
	}
	countS := 0
	l := j - i + 1
	for ; i <= j; i++ {
		if directions[i] == 'S' {
			countS++
		}
	}
	return l - countS
}

func countCollisions1(directions string) int {
	var r int
	var s bool
	res := 0
	for _, d := range directions {
		switch d {
		case 'L':
			if s {
				res += r + 1
				r = 0
			}
		case 'R':
			r++
			s = true
		default:
			res += r
			r = 0
			s = true
		}
	}
	return res
}

func countCollisions0(directions string) int {
	var r int
	var s bool
	res := 0
	for _, d := range directions {
		switch d {
		case 'L':
			if r > 0 || s {
				res += r + 1
				r = 0
				s = true
			}
		case 'R':
			r++
		default:
			res += r
			r = 0
			s = true
		}
	}
	return res
}
