package p3306

import "strings"

func countOfSubstrings(word string, k int) int64 {
	result := int64(0)
	minRight := 0
	maxRight := 0
	counter := newMyCounter()

	// Our window is the substring of word from left to right(exclusive)
	for left := 0; left <= len(word)-5; left++ {
		if 0 < left {
			counter.remove(word[left-1])
		}
		for minRight < len(word) && !counter.containAllVowels() {
			counter.add(word[minRight])
			minRight++
		}
		for minRight < len(word) && counter.consonantCnt < k {
			counter.add(word[minRight])
			minRight++
		}
		if minRight == len(word) && (!counter.containAllVowels() || counter.consonantCnt < k) {
			break
		}
		// If our code reaches here,
		// the substring word[left..<minRight] contains all the vowels and at least k consonants.

		if k < counter.consonantCnt {
			continue
		}
		if maxRight < minRight {
			maxRight = minRight
			for maxRight < len(word) && strings.IndexByte(counter.vowels, word[maxRight]) != -1 {
				maxRight++
			}
		}
		result += int64(maxRight - minRight + 1)
	}
	return result
}

type myCounter struct {
	vowels       string
	vowelFreq    []int
	consonantCnt int
}

func newMyCounter() *myCounter {
	return &myCounter{
		vowels:       "aeiou",
		vowelFreq:    make([]int, 5),
		consonantCnt: 0,
	}
}

func (counter *myCounter) add(c byte) {
	index := strings.IndexByte(counter.vowels, c)
	if index < 0 {
		counter.consonantCnt++
	} else {
		counter.vowelFreq[index]++
	}
}

func (counter *myCounter) remove(c byte) {
	index := strings.IndexByte(counter.vowels, c)
	if index < 0 {
		counter.consonantCnt--
	} else {
		counter.vowelFreq[index]--
	}
}

func (counter *myCounter) containAllVowels() bool {
	for _, freq := range counter.vowelFreq {
		if freq == 0 {
			return false
		}
	}
	return true
}

func countOfSubstrings0(word string, k int) int64 {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	n := len(word)
	count := 0

	for left := 0; left < n; left++ {
		vowelCount := make(map[byte]int)
		consonantCount := 0

		for right := left; right < n; right++ {
			char := word[right]
			if vowels[char] {
				vowelCount[char]++
			} else {
				consonantCount++
			}

			if consonantCount > k {
				break
			}

			if consonantCount == k && len(vowelCount) == 5 {
				count++
			}
		}
	}

	return int64(count)
}
