package main

import (
	"fmt"
	"math"
)

func main() {
	//asteroids := []int{5, 10, -5, -55, -55, -55}
	//asteroids := []int{-2, -2, 1, -2}
	//asteroids := []int{-2, -2, 1, -1}
	//asteroids := []int{10, 2, -5}
	//asteroids := []int{8, -8}
	//asteroids := []int{-2, 1, -1, -2}
	asteroids := []int{-2, 2, 1, -2}

	fmt.Println(asteroidCollision(asteroids))
}

func check(a int, b int) bool {
	return math.Abs(float64(a)) < math.Abs(float64(b))
}

func move(a []int, b, j int) {
	if check(a[j], b) {
		a[j] = b
	} else if a[j] == -1*b {
		j--
	}
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, len(asteroids))
	j := -1
	for i := 0; i < len(asteroids); i++ {
		curentAsteroid := asteroids[i]
		if j == -1 {
			j = 0
			stack[j] = curentAsteroid
		} else if stack[j] > 0 && stack[j] == -1*curentAsteroid {
			j--
			continue
		} else if stack[j] > 0 && curentAsteroid < 0 {
			if check(stack[j], curentAsteroid) {
				stack[j] = asteroids[i]
			}
			j--
			for ; j >= 0 && stack[j] > 0 && stack[j+1] < 0; j-- {
				if check(stack[j], stack[j+1]) {
					stack[j] = stack[j+1]
				} else if stack[j] == -1*curentAsteroid {
					j--
				}
			}
			j++
		} else {
			j++
			stack[j] = curentAsteroid
		}
	}
	return stack[:j+1]
}
