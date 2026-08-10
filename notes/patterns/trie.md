# Trie (Prefix Tree)

## When to Use

- **Prefix matching:** find all words with a given prefix
- **Word search:** check if a word exists in a dictionary
- **Autocomplete:** suggest completions for a prefix
- **XOR maximization:** find the number that maximizes XOR with a given value (bitwise trie)

Key signals: "words with common prefix," "dictionary lookup," "prefix search," "autocomplete"

## Core Template

```go
type TrieNode struct {
    children [26]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func Constructor() Trie {
    return Trie{root: &TrieNode{}}
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &TrieNode{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return false
        }
        node = node.children[idx]
    }
    return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return false
        }
        node = node.children[idx]
    }
    return true
}
```

## Key Insights

1. **Each node represents a prefix** — the path from root to any node spells a prefix
2. **`isEnd` marks complete words** — a prefix might exist without being a complete word
3. **Space: O(total characters)** — shared prefixes are stored once
4. **Time: O(L)** for insert/search/startswith where L = word length

## Variants

### Bitwise Trie (for XOR problems)
Instead of 26 children (a-z), each node has 2 children (0 and 1). Process bits from MSB to LSB:
- For XOR maximization: at each bit, try to go to the opposite bit if available
- Used for: maximum XOR pair, maximum XOR subarray

### Trie with DFS (word enumeration)
After inserting all words, use DFS from a prefix node to find all completions:
```go
func (t *Trie) FindWords(prefix string) []string {
    node := t.root
    for _, ch := range prefix {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return nil
        }
        node = node.children[idx]
    }
    var results []string
    var dfs func(n *TrieNode, current string)
    dfs = func(n *TrieNode, current string) {
        if n.isEnd {
            results = append(results, current)
        }
        for i, child := range n.children {
            if child != nil {
                dfs(child, current+string(rune('a'+i)))
            }
        }
    }
    dfs(node, prefix)
    return results
}
```

## Common Mistakes

1. **Forgetting `isEnd`** — without it, you can't distinguish between a prefix and a complete word
2. **Not allocating nodes lazily** — allocate `&TrieNode{}` only when needed, not for all 26 children upfront
3. **Using map instead of array** — for fixed alphabet (a-z), `[26]*TrieNode` is faster than `map[byte]*TrieNode`
4. **Off-by-one in bit operations** — for bitwise trie, start from the most significant bit (typically bit 30 or 31 for 32-bit integers)

## Classic Problems to Practice

| Problem | Key Idea |
|---|---|
| **LC 208** — Implement Trie | Core template: insert, search, startsWith |
| **LC 211** — Add and Search Word | Trie + DFS for wildcard matching |
| **LC 212** — Word Search II | Trie + backtracking on grid |
| **LC 421** — Maximum XOR of Two Numbers | Bitwise trie: greedily pick opposite bit |
| **LC 1268** — Search Suggestions System | Trie + DFS for autocomplete suggestions |