package p1813

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	arr1 := strings.Split(sentence1, " ")
	arr2 := strings.Split(sentence2, " ")
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] != arr2[0] {
			break
		}
		arr2 = arr2[1:]
		arr1 = arr1[1:]
	}
	for len(arr1) > 0 && len(arr2) > 0 {
		end1 := len(arr1) - 1
		end2 := len(arr2) - 1
		if arr1[end1] != arr2[end2] {
			break
		}
		arr2 = arr2[:end2]
		arr1 = arr1[:end1]
	}
	return len(arr1) == 0 || len(arr2) == 0
}
