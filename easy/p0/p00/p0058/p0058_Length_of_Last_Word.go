package p0058

import "fmt"

func lengthOfLastWord(s string) int {
	//Находим первый не пробел
	i := len(s) - 1
	for s[i] == 32 {
		i--
	}
	//Считаем кол-во символов до первого пробела
	c := i
	for i >= 0 && s[i] != 32 {
		i--
	}
	return c - i
}

func main() {
	//s := "   fly me   to   the moon  "
	s := "  moon  "
	fmt.Println(" "[0])
	fmt.Println(lengthOfLastWord(s))
}
