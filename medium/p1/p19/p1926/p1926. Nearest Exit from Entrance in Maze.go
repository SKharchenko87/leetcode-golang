package p1926

import (
	"math"
)

type kletka struct {
	up, right, down, left int
	exit                  bool
	countKletkaToExit     int
}

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	graph := map[int]kletka{}

	exitKeys := []int{}

	for i := 0; i < m*n; i++ {
		if maze[i/n][i%n] == '.' {
			cur := kletka{}
			cur.up, cur.down, cur.left, cur.right = i-n, i+n, i-1, i+1
			if i/n == 0 || i%n == 0 || (i+1)%n == 0 || i/n == m-1 {
				if i/n == 0 {
					cur.up = -1
				}
				if i%n == 0 {
					cur.left = -1
				}
				if (i+1)%n == 0 {
					cur.right = -1
				}
				if i/n == m-1 {
					cur.down = -1
				}

				if entrance[0]*n+entrance[1] == i {
					cur.exit = false
					cur.countKletkaToExit = math.MaxInt
				} else {
					cur.exit = true
					exitKeys = append(exitKeys, i)
					cur.countKletkaToExit = 0
				}
			} else {
				cur.countKletkaToExit = math.MaxInt
			}
			graph[i] = cur
		}
	}

	tmpKeys := []int{}
	curKeys := exitKeys
	flg := true
	var addKey = func(nextExit, k int) {
		x, ok := graph[k]
		if ok && x.countKletkaToExit > nextExit {
			tmpKeys = append(tmpKeys, k)
			x.countKletkaToExit = nextExit
			graph[k] = x
			flg = true
		}
	}
	for flg {
		flg = false
		for _, k := range curKeys {
			cur := graph[k]
			addKey(cur.countKletkaToExit+1, cur.up)
			addKey(cur.countKletkaToExit+1, cur.right)
			addKey(cur.countKletkaToExit+1, cur.down)
			addKey(cur.countKletkaToExit+1, cur.left)
		}
		curKeys = tmpKeys
		tmpKeys = []int{}
	}
	res := graph[entrance[0]*n+entrance[1]].countKletkaToExit
	if res == math.MaxInt {
		res = -1
	}
	return res
}
