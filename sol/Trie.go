package sol

type TrieNode struct {
	Children  map[rune]*TrieNode
	EndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			Children:  make(map[rune]*TrieNode),
			EndOfWord: false,
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, r := range word {
		if _, exists := cur.Children[r]; !exists {
			cur.Children[r] = &TrieNode{
				Children:  make(map[rune]*TrieNode),
				EndOfWord: false,
			}
		}
		cur = cur.Children[r]
	}
	cur.EndOfWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, r := range word {
		if _, exists := cur.Children[r]; !exists {
			return false
		}
		cur = cur.Children[r]
	}
	return cur.EndOfWord

}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, r := range prefix {
		if _, exists := cur.Children[r]; !exists {
			return false
		}
		cur = cur.Children[r]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
