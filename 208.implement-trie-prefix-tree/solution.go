package leetcode

type Trie struct {
	isLeaf   bool
	children map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isLeaf:   false,
		children: make(map[byte]*Trie, 0),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	_, exist := this.children[word[0]]
	if !exist {
		t := Constructor()
		this.children[word[0]] = &t
	}
	if len(word) == 1 {
		this.children[word[0]].isLeaf = true
	} else {
		this.children[word[0]].Insert(word[1:])
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return true
	}
	v, exist := this.children[word[0]]
	if !exist {
		return false
	}
	if len(word) == 1 {
		return v.isLeaf
	}
	return v.Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	v, exist := this.children[prefix[0]]
	if !exist {
		return false
	}
	return v.StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
