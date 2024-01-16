package p0380

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	keys     []int
	keyIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.keyIndex[val]; ok {
		return false
	}
	this.keys = append(this.keys, val)
	this.keyIndex[val] = len(this.keys) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.keyIndex[val]; !ok {
		return false
	}
	index := this.keyIndex[val]
	lastIndex := len(this.keys) - 1
	this.keys[index], this.keys[lastIndex] = this.keys[lastIndex], this.keys[index]
	this.keyIndex[this.keys[index]] = index
	this.keys = this.keys[:lastIndex]
	delete(this.keyIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.keys))
	return this.keys[index]
}

func run(commands []string, params [][]int) []interface{} {
	var cur RandomizedSet
	var res []interface{}
	for i, command := range commands {
		switch command {
		case "RandomizedSet":
			cur = Constructor()
			res = append(res, -1)
		case "insert":
			res = append(res, cur.Insert(params[i][0]))
			fmt.Println(cur.keys, cur.keyIndex)
		case "remove":
			res = append(res, cur.Remove(params[i][0]))
			fmt.Println(cur.keys, cur.keyIndex)
		case "getRandom":
			res = append(res, cur.GetRandom())
		}
	}
	return res
}
