package lc2213

import (
	"fmt"
)

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {

	type node struct {
		leftChar  byte
		rightChar byte
		pre       int
		suf       int
		iLen      int
		maxLen    int
		isEmpty   bool
	}

	type segtree struct {
		n    int
		tree []node
	}

	n := 1
	for n < len(s) {
		n *= 2
	}
	st := &segtree{
		n:    n,
		tree: make([]node, 2*n),
	}

	merge := func(left, right node) node {
		if left.isEmpty {
			return right
		}
		if right.isEmpty {
			return left
		}

		leftChar := left.leftChar
		rightChar := right.rightChar
		pre := left.pre
		suf := right.suf

		if left.rightChar == right.leftChar {
			if left.pre == left.iLen {
				pre = left.pre + right.pre
			}
			if right.suf == right.iLen {
				suf = right.suf + left.suf
			}
		}

		maxLen := max(left.maxLen, right.maxLen)
		if left.rightChar == right.leftChar {
			maxLen = max(maxLen, left.suf+right.pre)
		}

		return node{
			leftChar:  leftChar,
			rightChar: rightChar,
			pre:       pre,
			suf:       suf,
			iLen:      left.iLen + right.iLen,
			maxLen:    maxLen,
		}
	}

	build := func() {
		for i := 0; i < len(s); i++ {
			st.tree[n+i] = node{
				pre:       1,
				suf:       1,
				maxLen:    1,
				iLen:      1,
				leftChar:  s[i],
				rightChar: s[i],
			}
		}
		// mark padding leaves beyond len(s) as empty
		for i := len(s); i < n; i++ {
			st.tree[n+i] = node{isEmpty: true}
		}
		for i := n - 1; i > 0; i-- {
			st.tree[i] = merge(st.tree[i*2], st.tree[2*i+1])
		}
	}

	var update func(pos int, ch byte)
	update = func(pos int, ch byte) {
		i := st.n + pos

		st.tree[i] = node{
			leftChar:  ch,
			rightChar: ch,
			pre:       1,
			suf:       1,
			iLen:      1,
			maxLen:    1,
		}

		for i /= 2; i > 0; i /= 2 {
			st.tree[i] = merge(st.tree[i*2], st.tree[2*i+1])
		}
	}

	build()
	k := len(queryCharacters)
	res := make([]int, k)
	for i := range k {
		qc := queryCharacters[i]
		qi := queryIndices[i]

		update(qi, qc)
		res[i] = st.tree[1].maxLen
	}
	return res
}

func LC2213() {
	// s := "babacac"
	// queryCharacters := "bcb"
	// queryIndices := []int{1, 3, 3}

	// s := "abyzz"
	// queryCharacters := "aa"
	// queryIndices := []int{2, 1}

	s := "mm"
	queryCharacters := "bfviuwsr"
	queryIndices := []int{0, 0, 1, 0, 0, 1, 1, 0}

	// s := "dbgmcagale"
	// queryCharacters := "mf"
	// queryIndices := []int{5, 0}

	// s := "dbgmmmgale"
	// queryCharacters := "cf"
	// queryIndices := []int{4, 0} // [1,1]

	out := longestRepeating(s, queryCharacters, queryIndices)
	fmt.Println(out)
}
