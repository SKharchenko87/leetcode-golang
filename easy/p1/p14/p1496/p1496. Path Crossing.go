package main

func main() {
	path := "NESWW"
	print(isPathCrossing(path))
}

type point struct {
	x int
	y int
}

func isPathCrossing(path string) bool {
	x := 0
	y := 0
	p := point{x, y}
	m := map[point]bool{}
	m[p] = true
	for _, v := range path {
		switch v {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x--
		case 'W':
			x++
		}
		p = point{x, y}
		if m[p] {
			return true
		}
		m[p] = true
	}
	return false
}
