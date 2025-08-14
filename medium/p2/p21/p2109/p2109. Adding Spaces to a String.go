package p2109

import "strings"

func addSpaces(s string, spaces []int) string {
	bl := strings.Builder{}
	bl.Grow(len(s) + len(spaces))

	lastPosition := 0
	for _, space := range spaces {
		bl.WriteString(s[lastPosition:space])
		bl.WriteRune(' ')
		lastPosition = space
	}
	bl.WriteString(s[lastPosition:])
	return bl.String()
}

func addSpaces4(s string, spaces []int) string {
	bl := strings.Builder{}
	bl.Grow(len(s) + len(spaces))
	indexSpace := 0
	for index, c := range s {
		if indexSpace < len(spaces) && index == spaces[indexSpace] {
			bl.WriteRune(' ')
			indexSpace++
		}
		bl.WriteRune(c)
	}
	return bl.String()
}

func addSpaces3(s string, spaces []int) string {
	res := make([]byte, 0, len(s)+len(spaces))
	indexSpace := 0
	for index := 0; index < len(s); index++ {
		if indexSpace < len(spaces) && index == spaces[indexSpace] {
			res = append(res, ' ')
			indexSpace++
		}
		res = append(res, s[index])
	}
	return string(res)
}

func addSpaces2(s string, spaces []int) string {
	res := make([]byte, len(s)+len(spaces))
	indexSpace := 0
	i := 0
	for index := 0; index < len(s); index++ {
		if indexSpace < len(spaces) && index == spaces[indexSpace] {
			res[i] = ' '
			indexSpace++
			i++
		}
		res[i] = s[index]
		i++
	}
	return string(res)
}

func addSpaces1(s string, spaces []int) string {
	bl := strings.Builder{}
	indexSpace := 0
	for index, c := range s {
		if indexSpace < len(spaces) && index == spaces[indexSpace] {
			bl.WriteRune(' ')
			indexSpace++
		}
		bl.WriteRune(c)
	}
	return bl.String()
}

func addSpaces0(s string, spaces []int) string {
	res := make([]byte, len(s)+len(spaces))
	indexSpace := 0
	i := 0
	for index, c := range s {
		if indexSpace < len(spaces) && index == spaces[indexSpace] {
			res[i] = ' '
			indexSpace++
			i++
		}
		res[i] = byte(c)
		i++
	}
	return string(res)
}
