package p0744

import "slices"

func nextGreatestLetter(letters []byte, target byte) byte {
	index, _ := slices.BinarySearch(letters, target+1)
	return letters[index%len(letters)]
}
