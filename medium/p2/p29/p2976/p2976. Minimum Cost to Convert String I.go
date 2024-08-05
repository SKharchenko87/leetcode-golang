package p2976

import (
	"math"
	"strings"
)

//'a', 'b', 'c', 'c', 'e', 'd'}, changed: []byte{'b', 'c', 'b', 'e', 'b', 'e'
// a->b->c->e	0->1->2->4
// b->c->e 		1->2->4
// c->e c->b 	2->4 2->1
// e->b->c		4->1->2
// d->e			3->4

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	var res int64
	matrix := findMinPath(original, changed, cost)
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			if matrix[source[i]-'a'][target[i]-'a'] == 0 {
				return -1
			}
			res += int64(matrix[source[i]-'a'][target[i]-'a'])
		}
	}
	return res
}

func findMinPath(original []byte, changed []byte, cost []int) (res [26][26]int) {

	graph := map[byte]map[byte]int{}
	for i, b := range original {
		if res[b-'a'][changed[i]-'a'] > cost[i] || res[b-'a'][changed[i]-'a'] == 0 {
			res[b-'a'][changed[i]-'a'] = cost[i]
			if _, ok := graph[b-'a']; !ok {
				graph[b-'a'] = map[byte]int{}
			}
			graph[b-'a'][changed[i]-'a'] = cost[i]
		}
	}

	for b, m := range graph {
		for len(m) > 0 {
			newM := map[byte]int{}
			for b2, i := range m {
				for b3, i2 := range graph[b2] {
					if res[b][b3] > i+i2 || res[b][b3] == 0 {
						res[b][b3] = i + i2
						newM[b3] = i + i2
					}
				}
			}
			m = newM
		}
	}
	return res
}

type convert struct {
	from byte
	to   byte
}

func minimumCost0(source string, target string, original []byte, changed []byte, cost []int) int64 {
	m := map[convert]int{}
	mx := map[byte][]byte{}
	res := map[string]int{}
	for i, v := range original {
		res[string(v)]++
		if v != changed[i] {
			c := convert{from: v, to: changed[i]}
			m[c] = cost[i]
			if arr, ok := mx[v]; ok {
				mx[v] = append(arr, changed[i])
			} else {
				mx[v] = []byte{changed[i]}
			}

		}
	}

	rec(mx, &res, 1)

	for str, _ := range res {
		if len(str) > 1 {
			var minCost int
			if v, ok := m[convert{str[0], str[len(str)-1]}]; ok {
				minCost = v
			} else {
				minCost = math.MaxInt
			}
			cost := 0
			for i := 0; i < len(str)-1; i++ {
				cost += m[convert{str[i], str[i+1]}]
			}
			if cost < minCost {
				m[convert{str[0], str[len(str)-1]}] = cost
			}
		}
	}
	var total int64
	for i, src := range source {
		if byte(src) != target[i] {
			if c, ok := m[convert{byte(src), target[i]}]; ok {
				total += int64(c)
			} else {
				return -1
			}
		}
	}

	return total
}

func rec(mx map[byte][]byte, res *map[string]int, l int) {
	for v, _ := range *res {
		if len(v) == l {
			arrMX := mx[v[len(v)-1]]
			for i := 0; i < len(arrMX); i++ {
				_, ok := (*res)[v+string(arrMX[i])]
				if !strings.Contains(v, string(arrMX[i])) && !ok {
					(*res)[v+string(arrMX[i])]++
					rec(mx, res, l+1)
				}
			}
		}
	}
}
