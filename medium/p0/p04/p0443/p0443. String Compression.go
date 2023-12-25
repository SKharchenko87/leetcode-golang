package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compress(chars []byte) int {
	l := len(chars)
	counter := 1
	sb := strings.Builder{}
	for i := 1; i < l; i++ {
		if chars[i] != chars[i-1] {
			sb.WriteString(string(chars[i-1]))
			if counter != 1 {
				sb.WriteString(strconv.Itoa(counter))
			}
			counter = 1
		} else {
			counter++
		}
	}
	sb.WriteString(string(chars[l-1]))
	if counter != 1 {
		sb.WriteString(strconv.Itoa(counter))
	}
	str := sb.String()
	for i, b := range []byte(str) {
		chars[i] = b
	}
	return len(str)
}

func main() {
	//chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c'}
	chars := []byte{'a', 'a', 'b', 'b', 'b', 'c', 'c'}
	//chars := []byte{'a'}
	fmt.Println(compress(chars))
	fmt.Println(chars)
}
