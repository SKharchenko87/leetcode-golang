package p2751

import "sort"

type stack[T any] []T

func (s *stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack[T]) Len() int {
	return len(*s)
}

func (s *stack[T]) Peek() T {
	v := (*s)[len(*s)-1]
	return v
}

type mapIndex struct {
	index    int
	position int
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	mapIndexes := make([]mapIndex, len(positions))
	for i, position := range positions {
		mapIndexes[i] = mapIndex{i, position}
	}
	sort.Slice(mapIndexes, func(i, j int) bool {
		return mapIndexes[i].position < mapIndexes[j].position
	})
	var stackRobots stack[int]
loop:
	for i := 0; i < len(mapIndexes); i++ {
		if stackRobots.Len() == 0 {
			stackRobots.Push(i)
		} else {
			for stackRobots.Len() > 0 && directions[mapIndexes[stackRobots.Peek()].index] == 'R' && directions[mapIndexes[i].index] == 'L' {
				iIndex := mapIndexes[i].index
				lastIndex := mapIndexes[stackRobots.Peek()].index
				if healths[iIndex] == healths[lastIndex] {
					stackRobots.Pop()
					healths[iIndex] = 0
					healths[lastIndex] = 0
					continue loop
				} else if healths[iIndex] > healths[lastIndex] {
					stackRobots.Pop()
					healths[iIndex]--
					healths[lastIndex] = 0
				} else {
					healths[lastIndex]--
					healths[iIndex] = 0
					continue loop
				}
			}
			stackRobots.Push(i)
		}
	}

	res := make([]int, stackRobots.Len())
	j := 0
	for i := 0; i < len(healths); i++ {
		if healths[i] > 0 {
			res[j] = healths[i]
			j++
		}
	}
	return res
}

type robot struct {
	index     int
	position  int
	health    int
	direction uint8
}

func survivedRobotsHealths2(positions []int, healths []int, directions string) []int {
	robots := make([]robot, len(positions))
	for i, position := range positions {
		robots[i].index = i
		robots[i].position = position
		robots[i].health = healths[i]
		robots[i].direction = directions[i]
	}
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].position < robots[j].position
	})
	var stackRobots stack[robot]
loop:
	for i := 0; i < len(robots); i++ {
		if stackRobots.Len() == 0 {
			stackRobots.Push(robots[i])
		} else {
			for stackRobots.Len() > 0 && stackRobots.Peek().direction == 'R' && robots[i].direction == 'L' {
				r0 := stackRobots.Peek()
				if robots[i].health == r0.health {
					stackRobots.Pop()
					continue loop
				} else if robots[i].health > r0.health {
					stackRobots.Pop()
					robots[i].health--
				} else {
					r := stackRobots.Pop()
					r.health--
					stackRobots.Push(r)
					continue loop
				}
			}
			stackRobots.Push(robots[i])
		}
	}

	sort.Slice(stackRobots, func(i, j int) bool {
		return stackRobots[i].index < stackRobots[j].index
	})

	res := make([]int, stackRobots.Len())
	for i, r := range stackRobots {
		res[i] = r.health
	}
	return res
}

type stack1 []interface{}

func (s *stack1) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack1) Pop() interface{} {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func (s *stack1) Len() int {
	return len(*s)
}

func (s *stack1) Peek() interface{} {
	l := len(*s)
	v := (*s)[l-1]
	return v
}

func survivedRobotsHealths1(positions []int, healths []int, directions string) []int {
	robots := make([]robot, len(positions))
	for i, position := range positions {
		robots[i].index = i
		robots[i].position = position
		robots[i].health = healths[i]
		robots[i].direction = directions[i]
	}
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].position < robots[j].position
	})
	var x stack1
loop:
	for i := 0; i < len(robots); i++ {
		if x.Len() == 0 {
			x.Push(robots[i])
		} else {
			for x.Len() > 0 && x.Peek().(robot).direction == 'R' && robots[i].direction == 'L' {
				r0 := x.Peek().(robot)
				if robots[i].health == r0.health {
					x.Pop()
					continue loop
				} else if robots[i].health > r0.health {
					x.Pop()
					robots[i].health--
				} else {
					r := x.Pop().(robot)
					r.health--
					x.Push(r)
					continue loop
				}
			}
			x.Push(robots[i])
		}
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].(robot).index < x[j].(robot).index
	})

	res := make([]int, x.Len())
	for i, r := range x {
		res[i] = r.(robot).health
	}
	return res
}

func survivedRobotsHealths0(positions []int, healths []int, directions string) []int {
	robots := make([]robot, len(positions))
	for i, position := range positions {
		robots[i].index = i
		robots[i].position = position
		robots[i].health = healths[i]
		robots[i].direction = directions[i]
	}
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].position < robots[j].position
	})
	for i := 1; i < len(robots); i++ {
		if robots[i-1].direction == 'R' && robots[i].direction == 'L' {
			if robots[i].health == robots[i-1].health {
				robots = append(robots[:i-1], robots[i+1:]...)
				i--
				if i > 0 {
					i--
				}
			} else if robots[i].health > robots[i-1].health {
				robots[i].health--
				robots = append(robots[:i-1], robots[i:]...)
				i--
				if i > 0 {
					i--
				}
			} else {
				robots[i-1].health--
				robots = append(robots[:i], robots[i+1:]...)
				i--
			}
		}
	}
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].index < robots[j].index
	})
	res := make([]int, len(robots))
	for i, robot := range robots {
		res[i] = robot.health
	}
	return res
}
