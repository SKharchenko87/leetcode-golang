package p1233

import (
	"strings"
)

type rootFolder struct {
	children map[string]*rootFolder
	flg      bool
}

func removeSubfolders(folder []string) []string {
	trie := rootFolder{make(map[string]*rootFolder), false}
	for _, s := range folder {
		d := strings.Split(s, "/")
		parent := &trie
		for i := 1; i < len(d); i++ {
			if parent.flg {
				break
			}
			if _, ok := parent.children[d[i]]; !ok {
				parent.children[d[i]] = &rootFolder{children: make(map[string]*rootFolder)}
			}
			parent = parent.children[d[i]]
		}
		parent.flg = true

	}
	result := []string{}
	var dfs func(root rootFolder, str string)
	dfs = func(root rootFolder, str string) {
		for s, r := range root.children {
			if r.flg {
				result = append(result, str+"/"+s)
			} else {
				dfs(*r, str+"/"+s)
			}
		}
	}
	dfs(trie, "")
	return result
}
