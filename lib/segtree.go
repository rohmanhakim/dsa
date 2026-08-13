package lib

type SegTree struct {
	n    int
	tree []int
}

func NewSegTree(a []int) *SegTree {
	n := 1
	for n < len(a) {
		n *= 2
	}

	st := &SegTree{
		n:    n,
		tree: make([]int, 2*n),
	}

	// Leaves
	copy(st.tree[n:], a)

	// Build
	for i := n - 1; i > 0; i-- {
		st.tree[i] = st.tree[i*2] + st.tree[i*2+1]
	}

	return st
}

// Point update: a[pos] = value
func (st *SegTree) Set(pos, value int) {
	i := st.n + pos
	st.tree[i] = value

	for i /= 2; i > 0; i /= 2 {
		st.tree[i] = st.tree[i*2] + st.tree[i*2+1]
	}
}

// Range query: [l, r)
func (st *SegTree) Query(l, r int) int {
	res := 0

	l += st.n
	r += st.n

	for l < r {
		if l&1 == 1 {
			res += st.tree[l]
			l++
		}

		if r&1 == 1 {
			r--
			res += st.tree[r]
		}

		l /= 2
		r /= 2
	}

	return res
}
