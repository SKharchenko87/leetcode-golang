package p0091

import "fmt"

func numDecodings(s string) int {
	ls := len(s)
	m := map[byte]int{(s[ls-1] - '0'): 1}
	for i := ls - 2; i >= 0; i-- {
		fmt.Println(m)
		currentMap := map[byte]int{}
		for k, v := range m {
			cur := s[i] - '0'
			if cur == 0 {
				if k == 0 {
					return 0
				}
				if x, ok := currentMap[cur]; ok {
					currentMap[cur] = x + v
				} else {
					currentMap[cur] = v
				}
			}
			if k <= 9 {
				index := (cur)*10 + k
				if 0 < index && index < 27 && cur != 0 {
					if x, ok := currentMap[index]; ok {
						currentMap[index] = x + v
					} else {
						currentMap[index] = v
					}
				}
			}

			if 0 < cur && cur < 27 && k != 0 {
				if x, ok := currentMap[cur]; ok {
					currentMap[cur] = x + v
				} else {
					currentMap[cur] = v
				}

			}
		}
		m = currentMap
	}
	sum := 0
	for k, v := range m {
		if k == 0 {
			return 0
		}
		sum += v
	}
	fmt.Println(m)
	return sum
}

func NumDecodings(s string) int {
	return numDecodings(s)
}

func StrToSmart(s string) int {
	ls := len(s)
	m := map[byte]int{(s[ls-1] - '0'): 1}
	for i := ls - 2; i >= 0; i-- {
		fmt.Println(m)
		currentMap := map[byte]int{}
		for k, v := range m {
			cur := s[i] - '0'
			if cur == 0 {
				if k == 0 {
					return 0
				}
				if x, ok := currentMap[cur]; ok {
					currentMap[cur] = x + v
				} else {
					currentMap[cur] = v
				}
			}
			if k <= 9 {
				index := (cur)*10 + k
				if 0 < index && index < 27 && cur != 0 {
					if x, ok := currentMap[index]; ok {
						currentMap[index] = x + v
					} else {
						currentMap[index] = v
					}
				}
			}

			if 0 < cur && cur < 27 && k != 0 {
				if x, ok := currentMap[cur]; ok {
					currentMap[cur] = x + v
				} else {
					currentMap[cur] = v
				}

			}
		}
		m = currentMap
	}
	sum := 0
	for k, v := range m {
		if k == 0 {
			return 0
		}
		sum += v
	}
	fmt.Println(m)
	return sum
}
