package p1405

import (
	"container/heap"
	"sort"
)

type letterFreq struct {
	letter byte
	freq   int
}

type MaxHeap []letterFreq

func (lf *MaxHeap) Less(i, j int) bool {
	return (*lf)[i].freq > (*lf)[j].freq
}
func (lf *MaxHeap) Swap(i, j int) {
	(*lf)[i], (*lf)[j] = (*lf)[j], (*lf)[i]
}
func (lf *MaxHeap) Len() int {
	return len(*lf)
}
func (lf *MaxHeap) Pop() interface{} {
	old := *lf
	n := len(old)
	x := old[n-1]
	*lf = old[0 : n-1]
	return x
}
func (lf *MaxHeap) Push(x interface{}) {
	v := x.(letterFreq)
	*lf = append(*lf, v)
}

func longestDiverseString(a int, b int, c int) string {
	mh := MaxHeap{letterFreq{'a', a}, letterFreq{'b', b}, letterFreq{'c', c}}
	heap.Init(&mh)
	index := 0
	res := make([]byte, a+b+c)
	lastLetter, lastLetterCount := byte(0), 0
	for mh.Len() > 0 {

		x := heap.Pop(&mh).(letterFreq)
		if lastLetter != x.letter && x.freq > 0 {
			lastLetter = x.letter
			lastLetterCount = 0

			res[index] = x.letter
			index++
			lastLetterCount++
			x.freq--
			if x.freq > 0 {
				heap.Push(&mh, x)
			}
		} else if lastLetter == x.letter && lastLetterCount < 2 {
			res[index] = x.letter
			index++
			lastLetterCount++
			x.freq--
			if x.freq > 0 {
				heap.Push(&mh, x)
			}
		} else if lastLetter == x.letter && lastLetterCount >= 2 {
			if mh.Len() == 0 {
				break
			}
			y := heap.Pop(&mh).(letterFreq)
			heap.Push(&mh, x)
			x = y
			if x.freq > 0 {
				lastLetter = x.letter
				lastLetterCount = 0

				res[index] = x.letter
				index++
				lastLetterCount++
				x.freq--

				if x.freq > 0 {
					heap.Push(&mh, x)
				}
			} else {
				break
			}
		}
	}
	return string(res[:index])
}

func longestDiverseString0(a int, b int, c int) string {
	m := []letterFreq{{'a', a}, {'b', b}, {'c', c}}
	sort.Slice(m, func(i, j int) bool {
		return m[i].freq > m[j].freq
	})
	f12 := m[1].freq + m[2].freq
	if v := (f12)*2 + 2; m[0].freq > v {
		m[0].freq = v
	}
	index := 0
	res := make([]byte, m[0].freq+m[1].freq+m[2].freq)
	lastLetter, lastLetterCount := byte(0), 0
	for m[0].freq > 0 {
		if lastLetter != m[0].letter {
			lastLetter = m[0].letter
			lastLetterCount = 0

			res[index] = m[0].letter
			index++
			lastLetterCount++
			m[0].freq--

		} else if lastLetter == m[0].letter && lastLetterCount < 2 {
			res[index] = m[0].letter
			index++
			lastLetterCount++
			m[0].freq--
		} else if lastLetter == m[0].letter && lastLetterCount >= 2 {
			if m[1].freq > 0 {
				lastLetter = m[1].letter
				lastLetterCount = 0

				res[index] = m[1].letter
				index++
				lastLetterCount++
				m[1].freq--

			} else {
				break
			}
		}
		sort.Slice(m, func(i, j int) bool {
			return m[i].freq > m[j].freq
		})
	}
	return string(res)
}
