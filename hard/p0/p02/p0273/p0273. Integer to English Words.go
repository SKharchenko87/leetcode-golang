package p0273

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	res := ""
	dict := map[int]string{0: "", 1: "Thousand ", 2: "Million ", 3: "Billion "}
	k := 0
	m := map[int]string{1: "One ", 2: "Two ", 3: "Three ", 4: "Four ", 5: "Five ", 6: "Six ", 7: "Seven ", 8: "Eight ",
		9: "Nine ", 10: "Ten ", 11: "Eleven ", 12: "Twelve ", 13: "Thirteen ", 14: "Fourteen ", 15: "Fifteen ",
		16: "Sixteen ", 17: "Seventeen ", 18: "Eighteen ", 19: "Nineteen ", 20: "Twenty ", 30: "Thirty ", 40: "Forty ",
		50: "Fifty ", 60: "Sixty ", 70: "Seventy ", 80: "Eighty ", 90: "Ninety ",
	}

	for num > 0 {
		d99 := num % 100
		if d99 > 0 {
			res = dict[k] + res
			if d99 < 20 {
				res = m[d99] + res
			} else {
				res = m[d99/10*10] + m[d99%10] + res
			}
		}
		num /= 100
		d100 := num % 10
		if d100 > 0 {
			if d99 == 0 {
				res = dict[k] + res
			}
			res = m[d100] + "Hundred " + res
		}
		num /= 10
		k++
	}
	return strings.TrimSpace(res)
}

func numberToWords0(num int) string {
	if num == 0 {
		return "Zero"
	}
	res := ""
	dict := map[int]string{0: "", 1: "Thousand ", 2: "Million ", 3: "Billion "}
	k := 0

	for num > 0 {
		flg := true
		d := num % 100
		if d > 0 {
			res = dict[k] + res
			flg = false
			if d < 20 {
				res = getWord(d) + res
			} else {
				res = getWord(d/10*10) + getWord(d%10) + res
			}
		}
		num /= 100
		d = num % 10
		if d > 0 {
			if flg {
				res = dict[k] + res
			}
			res = getWord(d) + "Hundred " + res
		}
		num /= 10
		k++
	}
	return strings.TrimSpace(res)
}

func getWord(d int) string {
	switch d {
	case 1:
		return "One "
	case 2:
		return "Two "
	case 3:
		return "Three "
	case 4:
		return "Four "
	case 5:
		return "Five "
	case 6:
		return "Six "
	case 7:
		return "Seven "
	case 8:
		return "Eight "
	case 9:
		return "Nine "
	case 10:
		return "Ten "
	case 11:
		return "Eleven "
	case 12:
		return "Twelve "
	case 13:
		return "Thirteen "
	case 14:
		return "Fourteen "
	case 15:
		return "Fifteen "
	case 16:
		return "Sixteen "
	case 17:
		return "Seventeen "
	case 18:
		return "Eighteen "
	case 19:
		return "Nineteen "
	case 20:
		return "Twenty "
	case 30:
		return "Thirty "
	case 40:
		return "Forty "
	case 50:
		return "Fifty "
	case 60:
		return "Sixty "
	case 70:
		return "Seventy "
	case 80:
		return "Eighty "
	case 90:
		return "Ninety "
	}
	return ""
}
