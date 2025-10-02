package p3136

var vowels = [...]byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func isVowel(ch byte) bool {
	for _, vowel := range vowels {
		if ch == vowel {
			return true
		}
	}
	return false
}

func isConsonant(ch byte) bool {
	if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' {
		for _, vowel := range vowels {
			if ch == vowel {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func isDigit(ch byte) bool {
	if '0' <= ch && ch <= '9' {
		return true
	}
	return false
}

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	var vowel, consonant bool
	for i := range word {
		ch := word[i]
		if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9' {
			if isVowel(ch) {
				vowel = true
			} else if isConsonant(ch) {
				consonant = true
			}
		} else {
			return false
		}
	}
	return vowel && consonant
}

func is(word string) bool {
	cntCharacter := 0
	var vowel, consonant bool
	for i := range word {
		ch := word[i]
		if isDigit(ch) {
			cntCharacter++
		} else if isVowel(ch) {
			vowel = true
		} else if isConsonant(ch) {
			consonant = true
		} else {
			return false
		}
	}
	return cntCharacter >= 3 && vowel && consonant
}
