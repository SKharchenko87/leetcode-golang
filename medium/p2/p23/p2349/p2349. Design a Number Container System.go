package p2349

import "slices"

type NumberContainers struct {
	indexNumber   map[int]int
	numberIndexes map[int][]int
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int][]int{}}
}

func (this *NumberContainers) Change(index int, number int) {
	prevNumber, exists := this.indexNumber[index]
	if exists && prevNumber == number {
		return
	}

	if exists && prevNumber != number {
		i, _ := slices.BinarySearch(this.numberIndexes[prevNumber], index)
		this.numberIndexes[prevNumber] = slices.Delete(this.numberIndexes[prevNumber], i, i+1)
	}

	this.indexNumber[index] = number
	i, _ := slices.BinarySearch(this.numberIndexes[number], index)
	this.numberIndexes[number] = slices.Insert(this.numberIndexes[number], i, index)
}

func (this *NumberContainers) Find(number int) int {
	if len(this.numberIndexes[number]) == 0 {
		return -1
	}
	return this.numberIndexes[number][0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

func run(commands []string, params [][]int) []interface{} {
	var c NumberContainers
	result := make([]interface{}, len(commands))
	for i, command := range commands {
		switch command {
		case "NumberContainers":
			c = Constructor()
		case "find":
			result[i] = c.Find(params[i][0])
		case "change":
			c.Change(params[i][0], params[i][1])
		}
	}
	return result
}
