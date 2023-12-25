package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}
	m := make(map[int]int, len(deck))
	for _, v := range deck {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	c := m[deck[0]]
	if c%2 == 0 {
		for _, v := range m {
			if v%2 == 1 {
				return false
			}
		}
	} else {
		for _, v := range m {
			if c != v {
				mn := min(c, v)
				mx := max(c, v)
				if mx%mn != 0 {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	deck := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	fmt.Println(hasGroupsSizeX(deck))
}
