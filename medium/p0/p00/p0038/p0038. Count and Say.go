package main

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	res := strings.Builder{}
	str := countAndSay(n - 1)
	cur := str[0]
	counter := 1
	for i := 1; i < len(str); i++ {
		if cur == str[i] {
			counter++
		} else {
			//res.WriteString(fmt.Sprintf("%d%d", counter, cur-48))
			res.Write([]byte{byte(counter) + 48, cur})
			cur = str[i]
			counter = 1
		}
	}
	//res.WriteString(fmt.Sprintf("%d%d", counter, cur-48))
	res.Write([]byte{byte(counter) + 48, cur})
	return res.String()
}

func main() {
	fmt.Println(countAndSay(4))
}
