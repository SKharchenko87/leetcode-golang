package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	lenDigits := len(digits)
	res := []string{}

	m := map[byte]([]byte){
		50: {'a', 'b', 'c'},      //2
		51: {'d', 'e', 'f'},      //3
		52: {'g', 'h', 'i'},      //4
		53: {'j', 'k', 'l'},      //5
		54: {'m', 'n', 'o'},      //6
		55: {'p', 'q', 'r', 's'}, //7
		56: {'t', 'u', 'v'},      //8
		57: {'w', 'x', 'y', 'z'}, //9
	}

	sb := strings.Builder{}
	if lenDigits > 0 {
		arr0 := m[digits[0]]
		for _, v0 := range arr0 {
			if lenDigits > 1 {
				arr1 := m[digits[1]]
				for _, v1 := range arr1 {
					if lenDigits > 2 {
						arr2 := m[digits[2]]
						for _, v2 := range arr2 {
							if lenDigits > 3 {
								arr3 := m[digits[3]]
								for _, v3 := range arr3 {
									sb.Reset()
									sb.Write([]byte{v0, v1, v2, v3})
									res = append(res, sb.String())
								}
							} else {
								sb.Reset()
								sb.Write([]byte{v0, v1, v2})
								res = append(res, sb.String())
							}
						}
					} else {
						sb.Reset()
						sb.Write([]byte{v0, v1})
						res = append(res, sb.String())
					}
				}
			} else {
				sb.Reset()
				sb.Write([]byte{v0})
				res = append(res, sb.String())
			}
		}
	}
	return res
}

func main() {
	str := "23"
	fmt.Println(str[0] - 48)
	fmt.Println(str[1] - 48)

	fmt.Println(letterCombinations(str))
	//for _, c := range str {
	//	i, err := strconv.Atoi(string(c))
	//	if err == nil {
	//		fmt.Println(i)
	//	}
	//}

}
