package p3042

import "strings"

func countPrefixSuffixPairs(words []string) int {
	l := len(words)
	var count int
	for i := 1; i < l; i++ {
		lwi := len(words[i])
	LOOP:
		for j := 0; j < i; j++ {
			lwj := len(words[j])
			if lwj <= lwi {
				for k := 0; k < lwj; k++ {
					if words[i][k] != words[j][k] || words[i][lwi-k-1] != words[j][lwj-k-1] {
						continue LOOP
					}
				}
				count++
			}
		}
	}
	return count
}

func countPrefixSuffixPairs1(words []string) int {
	l := len(words)
	var count int
	for i := 1; i < l; i++ {
		lwi := len(words[i])
		for j := 0; j < i; j++ {
			lwj := len(words[j])
			if lwj < lwi && words[j] == words[i][:lwj] && words[j] == words[i][lwi-lwj:] {
				count++
			}
		}
	}
	return count
}

func countPrefixSuffixPairs0(words []string) int {
	l := len(words)
	var count int
	for i := 0; i < l-1; i++ {
		lwi := len(words[i])
	LOOP:
		for j := i + 1; j < l; j++ {
			lwj := len(words[j])
			if lwj >= lwi {
				for k := 0; k < lwi; k++ {
					if words[i][k] != words[j][k] || words[i][lwi-k-1] != words[j][lwj-k-1] {
						continue LOOP
					}
				}
				count++
			}
		}
	}
	return count
}

func countPrefixSuffixPairs3(words []string) int {
	l := len(words)
	var count int
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if strings.HasSuffix(words[j], words[i]) && strings.HasPrefix(words[j], words[i]) {
				count++
			}
		}
	}
	return count
}

func countPrefixSuffixPairs2(words []string) int {
	l := len(words)
	var count int
	for i := 0; i < l-1; i++ {
		lwi := len(words[i])
		for j := i + 1; j < l; j++ {
			lwj := len(words[j])
			if lwj >= lwi && words[i] == words[j][:lwi] && words[i] == words[j][lwj-lwi:] {
				count++
			}
		}
	}
	return count
}

type Node struct {
	links [26]*Node
}

func (this *Node) contains(c byte) bool {
	return this.links[c-'a'] != nil
}

func (this *Node) put(c byte, node *Node) {
	if !this.contains(c) {
		this.links = [26]*Node{}
	}
	this.links[c-'a'] = node
}

func (this *Node) next(c byte) *Node {
	return this.links[c-'a']
}

type Trie struct {
	root *Node
}

func (t *Trie) insert(word *string) {
	root := t.root
	for _, c := range *word {
		if !root.contains(byte(c)) {
			root.put(byte(c), &Node{})
		}
		root = root.next(byte(c))
	}
}

func (t *Trie) startWith(prefix *string) bool {
	root := t.root
	for _, c := range *prefix {
		if !root.contains(byte(c)) {
			return false
		}
		root = root.next(byte(c))
	}
	return true
}

func countPrefixSuffixPairs4(words []string) int {
	n := len(words)
	count := 0
	// Step 1: проходим по каждому слову
	for i := 0; i < n; i++ {
		var prefixTire, suffixTire Trie = Trie{&Node{}}, Trie{&Node{}}
		lwi := len(words[i])
		//Step 2: вставляем текущее слово в префексную Trie структуру
		prefixTire.insert(&words[i])

		//Step 3: вставляем текущее слово в суфиксную Trie структуру
		tmp := make([]byte, lwi)
		for k := 0; k < lwi; k++ {
			tmp[lwi-k-1] = words[i][k]
		}
		revWord := string(tmp)
		suffixTire.insert(&revWord)

		//Step 4: проходим по предыдущим словам
		for j := 0; j < i; j++ {
			lwj := len(words[j])
			//Step 5: Если длина j слова больше длины i слова, то пропускаем
			if lwj > lwi {
				continue
			}

			//Step 6: Получаем текущее и перевернутое слово
			tmp = make([]byte, lwj)
			for k := 0; k < lwj; k++ {
				tmp[lwj-k-1] = words[j][k]
			}
			revPrefixWord := string(tmp)

			//Step 7: проверяем префикс и суфикс
			if prefixTire.startWith(&words[j]) && suffixTire.startWith(&revPrefixWord) {
				count++
			}
		}
	}
	return count
}
