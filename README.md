# golang_implement_trie

A **[trie](https://en.wikipedia.org/wiki/Trie)** (pronounced as "try") or **prefix tree** is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

- `Trie()` Initializes the trie object.
- `void insert(String word)` Inserts the string `word` into the trie.
- `boolean search(String word)` Returns `true` if the string `word` is in the trie (i.e., was inserted before), and `false` otherwise.
- `boolean startsWith(String prefix)` Returns `true` if there is a previously inserted string `word` that has the prefix `prefix`, and `false` otherwise.

## Examples

**Example 1:**

```
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

```

**Constraints:**

- `1 <= word.length, prefix.length <= 2000`
- `word` and `prefix` consist only of lowercase English letters.
- At most $`3*10^4`$ calls **in total** will be made to `insert`, `search`, and `startsWith`.

## 解析

題目要我們實作一個可以儲存所有小寫英文字的一個 Trie結構

需要實作出以下 method:

1.  Constructor:  用來來建構 Trie
2.  insert(word string): 用來儲存 word
3. search(word string) bool: 用來察看是否有儲存過這個 word
4. startsWith(prefix string) bool:  用來察看是否有儲存過包含 prefix 的 word

首先知道 Trie 是一種特殊的 Tree

目前需要儲存所有字元是小寫英文 a-z 組成的字串

為了節省空間所以會是已字元為單位來做存儲

另外是可以透過把 字元當成每個 node 的 edge 上 key

如下圖：

![](https://i.imgur.com/JTyTsFg.png)

每個結點都可以透過 key 來找到下一個結點

這個結構可以使用 hashmap 來實作

而沒每個結點除了這個用來紀錄 child 對應的 hashmap 外

需要一個布林值來紀錄這個結點是否為最後一個結點

如上圖的 Trie 雖然有 appl 的 prefix 但是卻沒有 appl 這個字

## 程式碼

```go
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

```

## 困難點

1. 轉換紀錄 edge的值當作 index

## Solve Point

- [x]  Understand What problem need to solve
- [x]  Analysis complexity