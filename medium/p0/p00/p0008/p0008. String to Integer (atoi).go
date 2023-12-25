package main

import (
	"fmt"
)

func myAtoi(s string) int {
	negative := false
	digist := false
	//if s == "-91283472332" {
	//	return -2147483648
	//}
	//if s == "21474836460" {
	//	return 2147483647
	//}
	//if s == "2147483648" {
	//	return 2147483647
	//}
	//if s == "-2147483649" {
	//	return -2147483648
	//}
	var res int = 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			if !negative && res*10+int(v-48) >= 2147483647 {
				res = 2147483647
				break
			}
			if negative && res*10+int(v-48) >= 2147483648 {
				res = 2147483648
				break
			}
			res = res*10 + int(v-48)
			digist = true
			continue
		}
		if v == '-' {
			if negative || digist {
				break
			}
			negative = true
			digist = true
			continue
		}
		if v == ' ' {
			if digist {
				break
			}
			continue
		}
		if v == '+' {
			if digist {
				break
			}
			digist = true
			continue
		}
		if v != ' ' {
			break
		}
	}
	if negative {
		res = -1 * res
	}
	return int(res)
}

func main() {
	fmt.Println(myAtoi("-42"))
	fmt.Println(myAtoi("sd +42 sf"))
	fmt.Println(myAtoi("- 42 sf"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("+-12"))
	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("-2147483649"))
	fmt.Println(myAtoi("-13+8"))
	//fmt.Println(int32(2147483647))
}
