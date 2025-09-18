package p3408

import "container/heap"

type TaskManager struct {
	h tasksType
}

func Constructor(tasks [][]int) TaskManager {
	h := make([]*heapItem, 0, len(tasks))
	taskToId := map[int]int{}
	t := tasksType{h, taskToId}
	heap.Init(&t)
	for i := 0; i < len(tasks); i++ {
		x := heapItem{tasks[i][0], tasks[i][1], tasks[i][2], i}
		heap.Push(&t, &x)
	}

	return TaskManager{t}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	heap.Push(&this.h, &heapItem{userId, taskId, priority, this.h.Len()})
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	index := this.h.taskToId[taskId]
	this.h.items[index].priority = newPriority
	heap.Fix(&this.h, index)
}

func (this *TaskManager) Rmv(taskId int) {
	index := this.h.taskToId[taskId]
	heap.Remove(&this.h, index)
	delete(this.h.taskToId, taskId)
}

func (this *TaskManager) ExecTop() int {
	if this.h.Len() == 0 {
		return -1
	}
	tmp := heap.Pop(&this.h).(*heapItem)
	return tmp.userId
}

func run(commands []string, params []any) []any {
	result := make([]any, 0, len(commands))
	var tm TaskManager
	for i, command := range commands {
		switch command {
		case "TaskManager":
			tasks := params[i].([][]int)
			tm = Constructor(tasks)
			result = append(result, nil)
		case "add":
			addParams := params[i].([]int)
			tm.Add(addParams[0], addParams[1], addParams[2])
			result = append(result, nil)
		case "edit":
			editParams := params[i].([]int)
			tm.Edit(editParams[0], editParams[1])
			result = append(result, nil)
		case "execTop":
			result = append(result, tm.ExecTop())
		case "rmv":
			rmvParams := params[i].([]int)
			tm.Rmv(rmvParams[0])
			result = append(result, nil)
		}
	}
	return result
}

type heapItem struct {
	userId, taskId, priority int
	index                    int
}

type tasksType struct {
	items    []*heapItem
	taskToId map[int]int
}

func (this tasksType) Len() int {
	return len(this.items)
}
func (this tasksType) Less(i, j int) bool {
	if this.items[i].priority == this.items[j].priority {
		return this.items[i].taskId > this.items[j].taskId
	}
	return this.items[i].priority > this.items[j].priority
}

func (this tasksType) Swap(i, j int) {
	this.taskToId[this.items[i].taskId] = j
	this.taskToId[this.items[j].taskId] = i

	this.items[i], this.items[j] = this.items[j], this.items[i]
	this.items[i].index, this.items[j].index = j, i

}
func (this *tasksType) Push(x interface{}) {
	y := x.(*heapItem)
	this.taskToId[y.taskId] = y.index
	this.items = append(this.items, y)
}
func (this *tasksType) Pop() interface{} {
	old := this.items
	n := len(old)
	item := old[n-1]
	this.items = old[0 : n-1]
	return item
}
