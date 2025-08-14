package p0846

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	l := len(hand)
	if l%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	visit := hand[0] - 1
	for i := 0; i < l; i++ {
		if hand[i] == visit {
			continue
		}
		index := 1
		first := hand[i]
		hand[i] = visit
		for j := i + 1; j < l && index < groupSize; j++ {
			if hand[j] == visit {
				continue
			}
			// optimized
			if first+index < hand[j] {
				return false
			}
			if first+index == hand[j] {
				hand[j] = visit
				index++
			}
		}
		if index != groupSize {
			return false
		}
	}
	return true
}

func isNStraightHand1(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	for len(hand) > 0 {
		first := hand[0]
		hand = hand[1:]
		index := 1
		for j := 0; j < len(hand) && index < groupSize; j++ {
			if hand[j] > first+index {
				return false
			} else if hand[j] == first+index {
				hand = append(hand[:j], hand[j+1:]...)
				j--
				index++
			}
		}
		if index != groupSize {
			return false
		}
	}
	return true
}

func isNStraightHand0(hand []int, groupSize int) bool {
	l := len(hand)
	if l%groupSize != 0 {
		return false
	}
	sort.Ints(hand)

	for i := 0; i < l; i++ {
		if hand[i] == -1 {
			continue
		}
		k := 1
		tmp := hand[i]
		hand[i] = -1
		for j := i + 1; j < l && k < groupSize; j++ {
			if tmp+k == hand[j] && hand[j] != -1 {
				hand[j] = -1
				k++
			}
		}
		if k != groupSize {
			return false
		}
	}
	return true
}
