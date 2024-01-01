package p0208

import "fmt"

type Trie struct {
	m map[uint8]*Trie
}

func Constructor() Trie {
	newTrie := Trie{map[uint8]*Trie{'a': nil, 'b': nil, 'c': nil, 'd': nil, 'e': nil, 'f': nil, 'g': nil, 'h': nil, 'i': nil, 'j': nil, 'k': nil, 'l': nil, 'm': nil, 'n': nil, 'o': nil, 'p': nil, 'q': nil, 'r': nil, 's': nil, 't': nil, 'u': nil, 'v': nil, 'w': nil, 'x': nil, 'y': nil, 'z': nil}}
	return newTrie
}

func (this *Trie) Insert(word string) {
	trie := this
	for _, c := range word {
		if trie.m[uint8(c)] == nil {
			newTrie := Trie{map[uint8]*Trie{'a': nil, 'b': nil, 'c': nil, 'd': nil, 'e': nil, 'f': nil, 'g': nil, 'h': nil, 'i': nil, 'j': nil, 'k': nil, 'l': nil, 'm': nil, 'n': nil, 'o': nil, 'p': nil, 'q': nil, 'r': nil, 's': nil, 't': nil, 'u': nil, 'v': nil, 'w': nil, 'x': nil, 'y': nil, 'z': nil}}
			trie.m[uint8(c)] = &newTrie
		}
		trie = trie.m[uint8(c)]
	}
	trie.m[0] = nil
}

func (this *Trie) Search(word string) bool {
	trie := this
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i]
		if trie.m[c] == nil {
			return false
		}
		trie = trie.m[c]
	}
	if _, ok := trie.m[0]; ok {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	trie := this
	for _, c := range prefix {
		if trie.m[uint8(c)] == nil {
			return false
		}
		trie = trie.m[uint8(c)]
	}
	return true
}

func p0208(commands []string, args [][]string) [][]bool {
	obj := Constructor()
	l := len(commands)
	res := make([][]bool, l)
	for i := 1; i < l; i++ {
		switch commands[i] {
		case "insert":
			w := args[i][0]
			obj.Insert(w)
			res[i] = nil
		case "search":
			res[i] = []bool{obj.Search(args[i][0])}
		case "startsWith":
			res[i] = []bool{obj.StartsWith(args[i][0])}
		}
		fmt.Println(res)
	}
	return res
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
