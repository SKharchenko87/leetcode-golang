package p0726

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

type atom struct {
	name  string
	count int
}

var pow10 = []int{1, 10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 100_000_000}

func countOfAtoms(formula string) string {
	formulaItem := make([]*atom, len(formula))
	tmpName := make([]uint8, 2)
	cnt := 0
	coefficients := []int{}
	k := 1
	index := 0
	positionDigit := 0
	for i := len(formula) - 1; i >= 0; i-- {
		ch := formula[i]
		if (ch >= 'A' && ch <= 'Z') || ch == '(' || ch <= ')' {
			tmpName[0] = ch
			if cnt == 0 {
				cnt = 1
			}
			if ch == ')' {
				coefficients = append(coefficients, cnt)
				k *= cnt
			} else if ch == '(' {
				k /= coefficients[len(coefficients)-1]
				coefficients = coefficients[:len(coefficients)-1]
			} else {
				cnt *= k
				if tmpName[1] == 0 {
					tmpName = tmpName[:1]
				}
				a := &atom{string(tmpName), cnt}
				formulaItem[index] = a
				index++
			}

			tmpName = make([]uint8, 2)
			cnt = 0
			positionDigit = 0
		} else if ch >= 'a' && ch <= 'z' {
			tmpName[1] = ch
		} else if ch >= '0' && ch <= '9' {
			cnt = cnt + int(ch-'0')*pow10[positionDigit]
			positionDigit++
		}
	}

	formulaItem = formulaItem[:index]
	sort.Slice(formulaItem, func(i, j int) bool {
		return formulaItem[i].name < formulaItem[j].name
	})

	prevName := ""
	prevCount := 0
	var builder strings.Builder
	for i := 0; i < len(formulaItem); i++ {
		if prevName == formulaItem[i].name {
			prevCount += formulaItem[i].count
		} else {
			builder.WriteString(prevName)
			if prevCount > 1 {
				builder.WriteString(strconv.Itoa(prevCount))
			}
			prevName = formulaItem[i].name
			prevCount = formulaItem[i].count
		}
	}
	builder.WriteString(prevName)
	if prevCount > 1 {
		builder.WriteString(strconv.Itoa(prevCount))
	}
	return builder.String()
}

func countOfAtoms3(formula string) string {
	formulaItem := []*atom{}
	var atomStackPush = func(tmp string) {
		if len(tmp) > 0 {
			if tmp[0] == '(' || tmp[0] == ')' {
				v := atom{tmp, 1}
				formulaItem = append(formulaItem, &v)
			} else if tmp[0] >= 'A' && tmp[0] <= 'Z' {
				v := atom{tmp, 1}
				formulaItem = append(formulaItem, &v)
			} else if tmp[0] >= '0' && tmp[0] <= '9' {
				lastIndex := len(formulaItem) - 1
				formulaItem[lastIndex].count, _ = strconv.Atoi(tmp)
			}
		}
	}
	tmp := ""
	for _, ch := range formula {
		if ch >= 'A' && ch <= 'Z' {
			atomStackPush(tmp)
			tmp = string(ch)
		} else if ch >= 'a' && ch <= 'z' {
			tmp += string(ch)
		} else if ch >= '0' && ch <= '9' {
			if tmp[0] >= '0' && tmp[0] <= '9' {
				tmp += string(ch)
			} else {
				atomStackPush(tmp)
				tmp = string(ch)
			}
		} else {
			atomStackPush(tmp)
			tmp = string(ch)
		}
	}
	atomStackPush(tmp)

	koefs := []int{}
	k := 1
	for i := len(formulaItem) - 1; i >= 0; i-- {
		a := formulaItem[i]
		if a.name == ")" {
			koefs = append(koefs, a.count)
			k *= a.count
		} else if a.name == "(" {
			k /= koefs[len(koefs)-1]
			koefs = koefs[:len(koefs)-1]
		} else {
			a.count *= k
		}
	}

	sort.Slice(formulaItem, func(i, j int) bool {
		return formulaItem[i].name < formulaItem[j].name
	})

	res := ""
	prevName := ""
	prevCount := 0
	for i := 0; i < len(formulaItem); i++ {
		if formulaItem[i].name != "(" && formulaItem[i].name != ")" {
			if prevName == formulaItem[i].name {
				prevCount += formulaItem[i].count
			} else {
				res += prevName
				if prevCount > 1 {
					res += strconv.Itoa(prevCount)
				}
				prevName = formulaItem[i].name
				prevCount = formulaItem[i].count
			}
		}
	}
	res += prevName
	if prevCount > 1 {
		res += strconv.Itoa(prevCount)
	}
	return res
}

func countOfAtoms2(formula string) string {
	formulaItem := []*atom{}
	var atomStackPush = func(tmp string) {
		if len(tmp) > 0 {
			if tmp[0] == '(' || tmp[0] == ')' {
				v := atom{tmp, 1}
				formulaItem = append(formulaItem, &v)
			} else if tmp[0] >= 'A' && tmp[0] <= 'Z' {
				v := atom{tmp, 1}
				formulaItem = append(formulaItem, &v)
			} else if tmp[0] >= '0' && tmp[0] <= '9' {
				lastIndex := len(formulaItem) - 1
				formulaItem[lastIndex].count, _ = strconv.Atoi(tmp)
			}
		}
	}
	tmp := ""
	for _, ch := range formula {
		if ch >= 'A' && ch <= 'Z' {
			atomStackPush(tmp)
			tmp = string(ch)
		} else if ch >= 'a' && ch <= 'z' {
			tmp += string(ch)
		} else if ch >= '0' && ch <= '9' {
			if tmp[0] >= '0' && tmp[0] <= '9' {
				tmp += string(ch)
			} else {
				atomStackPush(tmp)
				tmp = string(ch)
			}
		} else {
			atomStackPush(tmp)
			tmp = string(ch)
		}
	}
	atomStackPush(tmp)

	koefs := []int{}
	k := 1
	for i := len(formulaItem) - 1; i >= 0; i-- {
		a := formulaItem[i]
		if a.name == ")" {
			koefs = append(koefs, a.count)
			k *= a.count
			formulaItem = append(formulaItem[:i], formulaItem[i+1:]...)
		} else if a.name == "(" {
			k /= koefs[len(koefs)-1]
			koefs = koefs[:len(koefs)-1]
			formulaItem = append(formulaItem[:i], formulaItem[i+1:]...)
		} else {
			a.count *= k
		}
	}

	sort.Slice(formulaItem, func(i, j int) bool {
		return formulaItem[i].name < formulaItem[j].name
	})

	res := ""
	prevName := ""
	prevCount := 0
	for i := 0; i < len(formulaItem); i++ {
		if prevName == formulaItem[i].name {
			prevCount += formulaItem[i].count
		} else {
			res += prevName
			if prevCount > 1 {
				res += strconv.Itoa(prevCount)
			}
			prevName = formulaItem[i].name
			prevCount = formulaItem[i].count
		}
	}
	res += prevName
	if prevCount > 1 {
		res += strconv.Itoa(prevCount)
	}
	return res
}

func countOfAtoms1(formula string) string {
	formulaItem := []atom{}
	var atomStackPush = func(tmp string) {
		if len(tmp) > 0 {
			if tmp[0] == '(' || tmp[0] == ')' {
				v := atom{tmp, 1}
				formulaItem = append(formulaItem, v)
			} else if tmp[0] >= 'A' && tmp[0] <= 'Z' {
				v := atom{tmp, 1}
				formulaItem = append(formulaItem, v)
			} else if tmp[0] >= '0' && tmp[0] <= '9' {
				lastIndex := len(formulaItem) - 1
				formulaItem[lastIndex].count, _ = strconv.Atoi(tmp)
			}
		}
	}
	tmp := ""
	for _, ch := range formula {
		if ch >= 'A' && ch <= 'Z' {
			atomStackPush(tmp)
			tmp = string(ch)
		} else if ch >= 'a' && ch <= 'z' {
			tmp += string(ch)
		} else if ch >= '0' && ch <= '9' {
			if tmp[0] >= '0' && tmp[0] <= '9' {
				tmp += string(ch)
			} else {
				atomStackPush(tmp)
				tmp = string(ch)
			}
		} else {
			atomStackPush(tmp)
			tmp = string(ch)
		}
	}
	atomStackPush(tmp)
	//fmt.Println(formulaItem)
	koefs := []int{}
	k := 1
	nameCount := map[string]int{}
	keys := []string{}
	for i := len(formulaItem) - 1; i >= 0; i-- {
		a := formulaItem[i]
		if a.name == ")" {
			koefs = append(koefs, a.count)
			k *= a.count
		} else if a.name == "(" {
			k /= koefs[len(koefs)-1]
			koefs = koefs[:len(koefs)-1]
		} else {
			a.count *= k
			if _, ok := nameCount[a.name]; ok {
				nameCount[a.name] += a.count
			} else {
				nameCount[a.name] = a.count
				keys = append(keys, a.name)
			}
		}
	}
	slices.Sort(keys)
	res := ""
	for _, key := range keys {
		res += key
		if nameCount[key] > 1 {
			res += strconv.Itoa(nameCount[key])
		}
	}
	return res
}

type stack []atom

func (s *stack) Push(v atom) {
	*s = append(*s, v)
}
func (s *stack) Pop() atom {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func (s *stack) Peek() *atom {
	l := len(*s)
	v := &((*s)[l-1])
	return v
}

func (s *stack) Empty() bool {
	return len(*s) == 0
}
func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) Print() {
	for _, x := range *s {
		print(x.name, x.count, " ")
	}
	println()
}

func countOfAtoms0(formula string) string {
	atomStack := stack{}
	var atomStackPush = func(tmp string) {
		if len(tmp) > 0 {
			if tmp[0] == '(' || tmp[0] == ')' {
				v := atom{tmp, 1}
				atomStack.Push(v)
			} else if tmp[0] >= 'A' && tmp[0] <= 'Z' {
				v := atom{tmp, 1}
				atomStack.Push(v)
			} else if tmp[0] >= '0' && tmp[0] <= '9' {
				v := atomStack.Peek()
				v.count, _ = strconv.Atoi(tmp)
			}
		}
	}
	tmp := ""
	for _, ch := range formula {
		if ch >= 'A' && ch <= 'Z' {
			atomStackPush(tmp)
			tmp = string(ch)
		} else if ch >= 'a' && ch <= 'z' {
			tmp += string(ch)
		} else if ch >= '0' && ch <= '9' {
			if tmp[0] >= '0' && tmp[0] <= '9' {
				tmp += string(ch)
			} else {
				atomStackPush(tmp)
				tmp = string(ch)
			}
		} else {
			atomStackPush(tmp)
			tmp = string(ch)
		}
	}
	atomStackPush(tmp)
	atomStack.Print() //ToDo
	koefs := []int{}
	k := 1
	nameCount := map[string]int{}
	keys := []string{}
	for atomStack.Len() > 0 {
		a := atomStack.Pop()
		if a.name == ")" {
			koefs = append(koefs, a.count)
			k *= a.count
		} else if a.name == "(" {
			k /= koefs[len(koefs)-1]
			koefs = koefs[:len(koefs)-1]
		} else {
			a.count *= k
			if _, ok := nameCount[a.name]; ok {
				nameCount[a.name] += a.count
			} else {
				nameCount[a.name] = a.count
				keys = append(keys, a.name)
			}
		}
	}
	slices.Sort(keys)
	res := ""
	for _, key := range keys {
		res += key
		if nameCount[key] > 1 {
			res += strconv.Itoa(nameCount[key])
		}
	}
	return res
}

//K4(ON(SO3)2)2
//K4(ONS2O6)2
//K4O2N2S22O12
//K4N2O14S4
