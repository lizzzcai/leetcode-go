package leetcode

type WordDictionary struct {
	next   map[byte]*WordDictionary
	isWord bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{make(map[byte]*WordDictionary), false}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		w := word[i]
		if m, ok := curr.next[w]; ok {
			curr = m
		} else {
			curr.next[w] = &WordDictionary{make(map[byte]*WordDictionary), false}
			curr = curr.next[w]
		}

	}
	curr.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return dfs(word, this)
}

func dfs(word string, node *WordDictionary) bool {
	if len(word) == 0 {
		return node.isWord
	}

	curr := word[0]
	if curr == '.' {
		for _, v := range node.next {
			if dfs(word[1:], v) {
				return true
			}
		}
	} else {
		if m, ok := node.next[curr]; ok {
			return dfs(word[1:], m)
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
