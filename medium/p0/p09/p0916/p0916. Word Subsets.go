package p0916

func wordSubsets(words1 []string, words2 []string) []string {
	hash2 := [26]byte{}
	tmp := [26]byte{}
	for _, word := range words2 {
		tmp = [26]byte{}
		for i := 0; i < len(word); i++ {
			tmp[word[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			if hash2[i] < tmp[i] {
				hash2[i] = tmp[i]
			}
		}
	}
	hash1 := [26]byte{}
	res := make([]string, 0, len(words1))
LOOP:
	for _, word := range words1 {
		hash1 = [26]byte{}
		for i := 0; i < len(word); i++ {
			hash1[word[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			if hash2[i] != 0 && hash1[i] < hash2[i] {
				continue LOOP
			}
		}
		res = append(res, word)
	}
	return res
}

func wordSubsets0(words1 []string, words2 []string) []string {
	hash2 := [26]byte{}
	var tmp [26]byte
	for _, word := range words2 {
		tmp = [26]byte{}
		for i := 0; i < len(word); i++ {
			tmp[word[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			if hash2[i] < tmp[i] {
				hash2[i] = tmp[i]
			}
		}
	}
	var hash1 [26]byte
	res := make([]string, 0, len(words1))
LOOP:
	for _, word := range words1 {
		hash1 = [26]byte{}
		for i := 0; i < len(word); i++ {
			hash1[word[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			if hash2[i] != 0 && hash1[i] < hash2[i] {
				continue LOOP
			}
		}
		res = append(res, word)
	}
	return res
}
