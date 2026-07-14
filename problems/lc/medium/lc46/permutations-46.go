package lc46

func permute(nums []int) [][]int {
	visited := make(map[int]bool)

	res := [][]int{}
	var findPermutation func(a int, path []int)

	findPermutation = func(a int, path []int) {
		visited[a] = true

		if len(path) == len(nums) {
			res = append(res, path)
			return
		}

		for _, b := range nums {
			if !visited[b] {
				path = append(path, b)
				findPermutation(b, path)
				path = path[:len(path)-1]
				visited[b] = false
			}
		}
	}

	for _, num := range nums {
		path := []int{}
		path = append(path, num)
		visited = make(map[int]bool)
		findPermutation(num, path)
	}
	return res
}
