package p1926

import "fmt"

type kletka struct {
	up, right, down, left int
	exit                  bool
}

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	graph := map[int]kletka{}

	for i := 0; i < m*n; i++ {
		if maze[i/n][i%n] == '.' {
			cur := kletka{}
			cur.up, cur.down, cur.left, cur.right = i-n, i+n, i-1, i+1
			if i == 0 || i%n == 0 || (i+1)%n == 0 || i/n == m-1 {
				if i == 0 {
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
				cur.exit = true
				if entrance[0]+entrance[1]*n == i {
					cur.exit = false
				}

			}
			graph[i] = cur
		}
	}

	fmt.Println(graph)

	//ToDo обход по спирали
	
	return 0
}
