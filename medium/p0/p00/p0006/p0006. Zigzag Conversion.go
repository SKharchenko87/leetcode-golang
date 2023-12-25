package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([]byte, len(s))
	if len(s) <= numRows {
		return s
	} else {
		k := 0
		for y := 0; y < numRows; y++ {
			for x := 0; x < len(s); x = x + (numRows-1)*2 {
				if y+x < len(s) {
					res[k] = s[y+x]
					//fmt.Println(string(res))
					k++
				}
				if 0 < y && y < numRows-1 {
					if x+(numRows-1)*2-y < len(s) {
						res[k] = s[x+(numRows-1)*2-y]
						//fmt.Println(string(res))
						k++
					}
				}
			}

		}
	}
	return string(res)
}

func main() {
	fmt.Println("AB")
	fmt.Println(convert("AB", 1))
	fmt.Println("AB")
}
