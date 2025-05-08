package main

import (
	"fmt"
	"strings"
)

const MaxN = 31

var RLE = [MaxN]string{}

func countAndSay(n int) string {
	return RLE[n]
}

func init() {
	RLE[1] = "1"
	for i := 2; i < MaxN; i++ {
		count := 1
		ch := RLE[i-1][0]
		tmp := make([]byte, 0, len(RLE[i-1])*4/3)
		for j := 1; j < len(RLE[i-1]); j++ {
			if ch == RLE[i-1][j] {
				count++
			} else {
				insert(&tmp, count, ch)
				count = 1
				ch = RLE[i-1][j]
			}
		}
		insert(&tmp, count, ch)
		RLE[i] = string(tmp)
	}
}

func insert(tmp *[]byte, count int, ch byte) {
	defer func() { *tmp = append(*tmp, ch) }()
	for count > 0 {
		defer func(c int) {
			*tmp = append(*tmp, byte(c)+'0')
		}(count % 10)
		count /= 10
	}
}

/*


















 */

func countAndSay0(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	res := strings.Builder{}
	str := countAndSay0(n - 1)
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
