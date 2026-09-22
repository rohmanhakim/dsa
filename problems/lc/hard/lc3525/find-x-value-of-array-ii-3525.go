package lc3525

func resultArray(nums []int, k int, queries [][]int) []int {

	type node struct {
		pre []int
		mul int64
	}

	type segtree struct {
		n    int
		tree []node
	}

	n := 1

	for n < len(nums) {
		n *= 2
	}

	st := &segtree{
		n:    n,
		tree: make([]node, 2*n),
	}

	merge := func(left, right node) node {
		spanningPre := make([]int, k)

		for y := range k {
			comb := (left.mul * int64(y)) % int64(k)
			spanningPre[comb] += right.pre[y]
		}

		mergedPre := make([]int, k)
		for i := 0; i < k; i++ {
			mergedPre[i] = left.pre[i] + spanningPre[i]
		}
		return node{
			pre: mergedPre,
			mul: (left.mul * right.mul) % int64(k),
		}
	}

	build := func() {
		for i := 0; i < len(nums); i++ {
			pre := make([]int, k)
			r := nums[i] % k
			pre[r]++

			st.tree[n+i] = node{
				pre: pre,
				mul: int64(nums[i] % k),
			}
		}

		for i := len(nums); i < n; i++ {
			st.tree[n+i] = node{
				pre: make([]int, k),
				mul: 1,
			}
		}

		for i := n - 1; i > 0; i-- {
			st.tree[i] = merge(st.tree[i*2], st.tree[2*i+1])
		}
	}

	var update func(index int, value int)
	update = func(index int, value int) {
		i := st.n + index

		pre := make([]int, k)
		r := value % k
		pre[r]++
		st.tree[i] = node{
			pre: pre,
			mul: int64(value % k),
		}

		for i /= 2; i > 0; i /= 2 {
			st.tree[i] = merge(st.tree[i*2], st.tree[2*i+1])
		}
	}

	var query func(nodeIdx, lo, hi, l, r int, result *node)
	query = func(nodeIdx, lo, hi, l, r int, result *node) {
		if hi < l || r < lo {
			return
		}

		if l <= lo && hi <= r {
			*result = merge(*result, st.tree[nodeIdx])
			return
		}

		mid := (lo + hi) / 2

		query(nodeIdx*2, lo, mid, l, r, result)
		query(nodeIdx*2+1, mid+1, hi, l, r, result)
	}

	build()
	ans := make([]int, len(queries))
	for i, q := range queries {
		update(q[0], q[1])
		res := node{
			pre: make([]int, k),
			mul: 1,
		}
		query(1, 0, n-1, q[2], len(nums)-1, &res)
		ans[i] = res.pre[q[3]]
	}

	return ans
}
