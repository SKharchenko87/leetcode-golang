package p1750

func minimumLength(s string) int {
	indexLeft, indexRight := 0, len(s)-1
	for indexLeft < indexRight {
		curLeft, curRight := s[indexLeft], s[indexRight]
		if curLeft != curRight {
			break
		}
		for indexLeft < len(s) && s[indexLeft] == curLeft {
			indexLeft++
		}
		for indexRight >= 0 && s[indexRight] == curRight {
			indexRight--
		}
	}
	if indexRight-indexLeft >= 0 {
		return indexRight - indexLeft + 1
	} else {
		return 0
	}
}
