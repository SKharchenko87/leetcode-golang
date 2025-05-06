package p2071

import (
	"fmt"
	"sort"
)

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	//ToDo
	lengthTasks := len(tasks)
	lengthWorkers := len(workers)
	newTasks := make([]int, 0, lengthTasks+lengthWorkers)
	newWorkers := make([]int, 0, lengthTasks+lengthWorkers)
	i := 0
	j := 0
	sort.Ints(tasks)
	sort.Ints(workers)
	res := 0
	for index := 0; index < lengthTasks+lengthWorkers; index++ {

		if i < lengthTasks && j < lengthWorkers && tasks[i] < workers[j] || j >= lengthWorkers {
			fmt.Println(tasks[i])
			newTasks = append(newTasks, tasks[i])
			newWorkers = append(newWorkers, -1)
			i++
			continue
		} else if i < lengthTasks && j < lengthWorkers && tasks[i] > workers[j] || i >= lengthTasks {
			fmt.Println(workers[j])
			newTasks = append(newTasks, -1)
			newWorkers = append(newWorkers, workers[j])
			j++
			continue
		}
		i++
		j++
		index++
		res++
	}
	index := 0
	for _, worker := range newWorkers {
		if worker != -1 {
			for ; newTasks[index] == -1 && index < lengthTasks+lengthWorkers; index++ {
			}
			if index < lengthTasks+lengthWorkers {
				if worker+strength >= newTasks[index] {
					res++
					pills--
					if pills == 0 {
						return res
					}
				}
			}
		}
	}
	return res
}
