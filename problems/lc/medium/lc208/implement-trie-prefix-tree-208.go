package lc208

import "fmt"

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.


Example 1:

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


Constraints:

1 <= word.length, prefix.length <= 2000
word and prefix consist only of lowercase English letters.
At most 3 * 10^4 calls in total will be made to insert, search, and startsWith.
*/

type Node struct {
	children []*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	root := Node{
		children: make([]*Node, 26),
	}
	return Trie{
		root: &root,
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root

	for _, ch := range word {
		if curr.children[ch-'a'] == nil {
			curr.children[ch-'a'] = &Node{
				children: make([]*Node, 26),
			}
		}
		curr = curr.children[ch-'a']
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root

	for _, ch := range word {
		if curr.children[ch-'a'] == nil {
			return false
		} else {
			curr = curr.children[ch-'a']
		}
	}

	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root

	for _, ch := range prefix {
		if curr.children[ch-'a'] == nil {
			return false
		}
		curr = curr.children[ch-'a']
	}
	return true
}

func ImplementTriePrefixTree() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
