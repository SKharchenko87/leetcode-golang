package p0502

import "container/heap"
import "sort"

type MaxHeap [][2]int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i][1] > m[j][1]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.([2]int))
}

func (m *MaxHeap) Pop() any {
	if len(*m) == 0 {
		return 0
	}
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var partition = func(low, high int) ([]int, int) {
		pivot := capital[high]
		i := low
		for j := low; j < high; j++ {
			if capital[j] < pivot {
				capital[i], capital[j] = capital[j], capital[i]
				profits[i], profits[j] = profits[j], profits[i]
				i++
			}
		}
		capital[i], capital[high] = capital[high], capital[i]
		profits[i], profits[high] = profits[high], profits[i]
		return capital, i
	}

	var quickSort func(low, high int) []int
	quickSort = func(low, high int) []int {
		if low < high {
			var p int
			capital, p = partition(low, high)
			capital = quickSort(low, p-1)
			capital = quickSort(p+1, high)
		}
		return capital
	}

	var quickSortStart = func() {
		quickSort(0, len(capital)-1)
	}
	quickSortStart()

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for j := 0; j < len(capital); j++ {
		if w >= capital[j] {
			heap.Push(maxHeap, [2]int{capital[j], profits[j]})
		} else if maxHeap.Len() > 0 && k > 0 {
			w += heap.Pop(maxHeap).([2]int)[1]
			k--
			if k == 0 {
				return w
			}
			j--
		}
	}
	for maxHeap.Len() > 0 && k > 0 {
		w += heap.Pop(maxHeap).([2]int)[1]
		k--
	}
	return w
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func findMaximizedCapital2(k int, w int, profits []int, capital []int) int {
	cm := make([][2]int, len(capital))
	for i := range capital {
		cm[i][0] = capital[i]
		cm[i][1] = profits[i]
	}
	sort.Slice(cm, func(i, j int) bool {
		if cm[i][1] == cm[j][1] {
			return cm[i][0] < cm[i][0]
		}
		return cm[i][1] > cm[j][1]
	})

	for i := 0; i < k; i++ {
		bestProfit := -1
		index := 0
		for j := index; j < len(cm); j++ {
			if cm[j][0] <= w {
				if bestProfit < cm[j][1] {
					bestProfit = cm[j][1]
					index = j
				}
				bestProfit = cm[j][1]
				index = j
				break
			}
		}
		if bestProfit == -1 {
			return w
		} else {
			//println(bestCapital, bestProfit)
			w += bestProfit
			cm = append(cm[:index], cm[index+1:]...)
			//cm[index][1] = -1
		}
	}
	return w
}

func findMaximizedCapital1(k int, w int, profits []int, capital []int) int {
	cm := make([][2]int, len(capital))
	for i := range capital {
		cm[i][0] = capital[i]
		cm[i][1] = profits[i]
	}
	sort.Slice(cm, func(i, j int) bool {
		if cm[i][0] == cm[j][0] {
			return cm[i][1] > cm[i][1]
		}
		return cm[i][0] < cm[j][0]
	})

	for i := 0; i < k; i++ {
		bestProfit := -1
		index := 0
		for j := index; j < len(cm); j++ {
			if cm[j][0] > w {
				break
			}
			if bestProfit < cm[j][1] {
				bestProfit = cm[j][1]
				index = j
			}
		}
		if bestProfit == -1 {
			return w
		} else {
			//println(bestCapital, bestProfit)
			w += bestProfit
			cm[index][1] = -1
		}
	}
	return w
}

func findMaximizedCapital0(k int, w int, profits []int, capital []int) int {
	capitals := make(map[int][]int)
	key := []int{}
	for i, c := range capital {
		if _, ok := capitals[c]; !ok {
			capitals[c] = []int{profits[i]}
			key = append(key, c)
		} else {
			capitals[c] = append(capitals[c], profits[i])
		}
	}
	sort.Ints(key)

	for i := 0; i < k; i++ {
		bestProfit := -1
		bestCapital := -1
		bestCapitalIndex := -1
		for _, j := range key {
			if j > w {
				break
			}
			for m := 0; m < len(capitals[j]); m++ {
				if capitals[j][m] > bestProfit {
					bestProfit = capitals[j][m]
					bestCapital = j
					bestCapitalIndex = m
				}
			}
		}
		if bestCapitalIndex == -1 {
			return w
		} else {
			println(bestCapital, bestProfit)
			w += bestProfit
			capitals[bestCapital][bestCapitalIndex] = -1
		}
	}
	return w
}
