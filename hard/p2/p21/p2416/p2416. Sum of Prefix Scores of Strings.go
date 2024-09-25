package p2416

type NodeTree struct {
	children [26]*NodeTree
	cnt      int
}

func (nt *NodeTree) InsertWord(word string) {
	currentTree := nt
	for _, ch := range word {
		ch = ch - 'a'
		if currentTree.children[ch] == nil {
			currentTree.children[ch] = &NodeTree{}
		}
		currentTree = currentTree.children[ch]
		currentTree.cnt++
	}
}

func (nt *NodeTree) SumScore(word string) (res int) {
	currentTree := nt
	l := len(word)
	for i := 0; i < l && currentTree != nil; i++ {
		ch := word[i] - 'a'
		res += currentTree.children[int32(ch)].cnt
		currentTree = currentTree.children[int32(ch)]
	}
	return res
}

func sumPrefixScores(words []string) []int {
	tree := &NodeTree{}
	for _, word := range words {
		tree.InsertWord(word)
	}
	res := make([]int, len(words))
	for i, word := range words {
		res[i] = tree.SumScore(word)
	}
	return res
}

type NodeTree1 struct {
	children map[int32]*NodeTree1
	cnt      int
}

func (nt *NodeTree1) InsertWord(word string) {
	currentTree := nt
	for _, ch := range word {
		if currentTree.children == nil {
			currentTree.children = make(map[int32]*NodeTree1)
		}
		if _, ok := currentTree.children[ch]; !ok {
			currentTree.children[ch] = &NodeTree1{}
		}
		currentTree = currentTree.children[ch]
		currentTree.cnt++
	}
}

func (nt *NodeTree1) SumScore(word string) (res int) {
	currentTree := nt
	l := len(word)
	for i := 0; i < l && currentTree != nil; i++ {
		ch := word[i]
		res += currentTree.children[int32(ch)].cnt
		currentTree = currentTree.children[int32(ch)]
	}
	return res
}

func sumPrefixScores1(words []string) []int {
	tree := &NodeTree1{children: make(map[int32]*NodeTree1), cnt: 0}
	for _, word := range words {
		tree.InsertWord(word)
	}
	res := make([]int, len(words))
	for i, word := range words {
		res[i] = tree.SumScore(word)
	}
	return res
}

func sumPrefixScores0(words []string) []int {
	tree := &NodeTree1{children: make(map[int32]*NodeTree1), cnt: 0}
	for _, word := range words {
		currentTree := tree
		for _, ch := range word {
			if currentTree.children == nil {
				currentTree.children = make(map[int32]*NodeTree1)
			}
			if _, ok := currentTree.children[ch]; !ok {
				currentTree.children[ch] = &NodeTree1{}
			}
			currentTree = currentTree.children[ch]
			currentTree.cnt++
		}
	}
	res := make([]int, len(words))
	for i, word := range words {
		currentTree := tree
		for _, ch := range word {
			res[i] += currentTree.children[ch].cnt
			currentTree = currentTree.children[ch]
		}
	}
	return res
}
