package p3606

import (
	"regexp"
	"sort"
	"unicode"
)

const (
	electronics byte = iota
	grocery
	pharmacy
	restaurant
	other
)

type validCoupons struct {
	code  *[]string
	order *[]byte
}

func (v validCoupons) Len() int {
	return len(*v.order)
}

func (v validCoupons) Less(i, j int) bool {
	if (*v.order)[i] == (*v.order)[j] {
		return (*v.code)[i] < (*v.code)[j]
	}
	return (*v.order)[i] < (*v.order)[j]
}

func (v validCoupons) Swap(i, j int) {
	(*v.order)[i], (*v.order)[j] = (*v.order)[j], (*v.order)[i]
	(*v.code)[i], (*v.code)[j] = (*v.code)[j], (*v.code)[i]
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	l := len(code)
	order := make([]byte, 0, l)
	newLen := 0
LOOP:
	for i := 0; i < l; i++ {
		if isActive[i] {
			if len(code[i]) == 0 {
				order = append(order, other)
				continue LOOP
			}
			for j := 0; j < len(code[i]); j++ {
				if unicode.IsLetter(rune(code[i][j])) || unicode.IsNumber(rune(code[i][j])) || code[i][j] == '_' {
					continue
				}
				order = append(order, other)
				continue LOOP
			}

			switch businessLine[i] {
			case "electronics":
				order = append(order, electronics)
			case "grocery":
				order = append(order, grocery)
			case "pharmacy":
				order = append(order, pharmacy)
			case "restaurant":
				order = append(order, restaurant)
			default:
				order = append(order, other)
				continue
			}
			newLen++
		} else {
			order = append(order, other)
		}
	}

	vc := validCoupons{&code, &order}
	sort.Sort(vc)

	return code[:newLen]
}

func validateCoupons2(code []string, businessLine []string, isActive []bool) []string {
	l := len(code)
	electronicsRes := make([]string, 0, l/4)
	groceryRes := make([]string, 0, l/4)
	pharmacyRes := make([]string, 0, l/4)
	restaurantRes := make([]string, 0, l/4)

	for i := 0; i < l; i++ {
		if isActive[i] {
			isValidCode, _ := regexp.MatchString(`^[A-Za-z0-9_]+$`, code[i])
			if isValidCode {
				switch businessLine[i] {
				case "electronics":
					electronicsRes = append(electronicsRes, code[i])
				case "grocery":
					groceryRes = append(groceryRes, code[i])
				case "pharmacy":
					pharmacyRes = append(pharmacyRes, code[i])
				case "restaurant":
					restaurantRes = append(restaurantRes, code[i])
				}
			}
		}
	}

	sort.Strings(electronicsRes)
	sort.Strings(groceryRes)
	sort.Strings(pharmacyRes)
	sort.Strings(restaurantRes)

	res := make([]string, 0, len(electronicsRes)+len(groceryRes)+len(pharmacyRes)+len(restaurantRes))
	res = append(res, electronicsRes...)
	res = append(res, groceryRes...)
	res = append(res, pharmacyRes...)
	res = append(res, restaurantRes...)
	return res
}

func validateCoupons1(code []string, businessLine []string, isActive []bool) []string {
	l := len(code)
	order := make([]byte, 0, l)
	newLen := 0
	for i := 0; i < l; i++ {
		if isActive[i] {
			isValidCode, _ := regexp.MatchString(`^[A-Za-z0-9_]+$`, code[i])
			if isValidCode {
				switch businessLine[i] {
				case "electronics":
					order = append(order, electronics)
				case "grocery":
					order = append(order, grocery)
				case "pharmacy":
					order = append(order, pharmacy)
				case "restaurant":
					order = append(order, restaurant)
				default:
					order = append(order, other)
					continue
				}
				newLen++
				continue
			}
		}
		order = append(order, other)
	}

	vc := validCoupons{&code, &order}
	sort.Sort(vc)

	return code[:newLen]
}

func validateCoupons0(code []string, businessLine []string, isActive []bool) []string {
	l := len(code)
	res := make([]string, 0, l)
	order := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		if isActive[i] {
			isValidCode, _ := regexp.MatchString(`^[A-Za-z0-9_]+$`, code[i])
			if isValidCode {
				switch businessLine[i] {
				case "electronics":
					res = append(res, code[i])
					order = append(order, electronics)
				case "grocery":
					res = append(res, code[i])
					order = append(order, grocery)
				case "pharmacy":
					res = append(res, code[i])
					order = append(order, pharmacy)
				case "restaurant":
					res = append(res, code[i])
					order = append(order, restaurant)
				}
			}
		}
	}

	vc := validCoupons{&res, &order}
	sort.Sort(vc)

	return res
}
