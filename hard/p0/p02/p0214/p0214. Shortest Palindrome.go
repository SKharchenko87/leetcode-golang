package p0214

import "strings"

func shortestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	hash := make([]int, l)
	hash[0] = int(s[0] - 'a')
	for i := 1; i < l; i++ {
		hash[i] = hash[i-1] + int(s[i]-'a')
	}

	for i := l - 1; i >= 0; i-- {
		var leftHash int
		if i%2 == 0 {
			leftHash = hash[i] - hash[i/2]
		} else {
			leftHash = hash[i] / 2
		}

		rightHash := hash[(i-1)/2]
		if leftHash == rightHash && s[:i/2+1] == stringRevertSlice(s[(i+1)/2:i+1]) {
			return stringRevertSlice(s[i+1:]) + s
		}

	}
	return stringRevertSlice(s[1:]) + s
}

func shortestPalindrome3(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	hash := make([]int, l)
	hash[0] = int(s[0] - 'a')
	indexPalindrome := 0
	for i := 1; i < l; i++ {
		hash[i] = hash[i-1] + int(s[i]-'a')
		if i%2 == 0 {
			if hash[i]-hash[i/2] == hash[(i-1)/2] && s[:i/2+1] == stringRevertSlice(s[(i+1)/2:i+1]) {
				indexPalindrome = i
			}
		} else {
			if hash[i]/2 == hash[(i-1)/2] && s[:i/2+1] == stringRevertSlice(s[(i+1)/2:i+1]) {
				indexPalindrome = i
			}
		}
	}
	return stringRevertSlice(s[indexPalindrome+1:]) + s
}

func stringRevertDeffer(in string) (out string) {
	sb := strings.Builder{}
	defer func() { out = sb.String() }()
	length := len(in)
	for i := 0; i < length; i++ {
		defer sb.WriteByte(in[i])
	}
	return
}

func stringRevertSlice(in string) (out string) {
	length := len(in)
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[length-1-i] = in[i]
	}
	return string(res)
}

func shortestPalindrome2(s string) string {
	l := len(s)
	sReverse := stringRevertSlice(s)
	for i := l - 1; i >= 0; i-- {
		if sReverse[l-1-i:] == s[:i+1] {
			return sReverse[:l-1-i] + s
		}
	}
	return ""
}

func shortestPalindrome1(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	start := 0
LOOP:
	for i := l - 1; i >= 0; i-- {
		if s[i] == s[0] {
			j := 1
			for ; j < i-j && s[j] == s[i-j]; j++ {

			}
			if 2*j >= i {
				start = i
				break LOOP
			}
		}
	}
	res := make([]byte, l-start-1)
	lenRes := len(res)
	for i := 0; i < lenRes; i++ {
		res[i] = s[l-i-1]
	}
	return string(res) + s
}

func shortestPalindrome0(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	start := -1
LOOP:
	for i := l - 1; i >= 0; i-- {
		if s[i] == s[0] {
			j := 1
			for ; j < i-j && s[j] == s[i-j]; j++ {

			}
			if 2*j >= i {
				start = i
				break LOOP
			}
		}
	}
	var res []byte
	if start == -1 {
		res = make([]byte, 2*l-1)
	} else {
		res = make([]byte, l+(l-start-1))
	}
	lenRes := len(res)
	for i := 0; i <= lenRes/2; i++ {
		res[i] = s[l-i-1]
		res[lenRes-i-1] = s[l-i-1]
	}
	return string(res)
}

//aacecaaa
