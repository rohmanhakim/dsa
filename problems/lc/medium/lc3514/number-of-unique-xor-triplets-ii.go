package lc3514

/*
func uniqueXorTriplets(nums []int) int {
	uniqueMap := make(map[int]bool)
	uniques := []int{}
	for _, num := range nums {
		if !uniqueMap[num] {
			uniques = append(uniques, num)
		}
		uniqueMap[num] = true
	}

	path := []int{}
	var dfs func(idx int)
	dfs = func(idx int) {

		if len(path) == 3 {

			res := uniques[path[0]] ^ uniques[path[1]] ^ uniques[path[2]]
			// fmt.Println(uniques[path[0]], uniques[path[1]], uniques[path[2]], res)

			uniqueMap[res] = true
			return
		}

		for i := idx; i < len(uniques); i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	// fmt.Println(uniques)

	return len(uniqueMap)
}
*/

func uniqueXorTriplets(nums []int) int {

	one := make([]bool, 1501)
	two := make([]bool, 2048)

	for _, num := range nums {
		one[num] = true
	}

	for _, num := range nums {
		for i, val := range one {
			if val {
				two[num^i] = true
			}
		}
	}

	three := make([]bool, 2048)
	for _, num := range nums {
		for i, val := range two {
			if val {
				three[num^i] = true
			}
		}
	}

	res := 0
	for _, val := range three {
		if val {
			res++
		}
	}

	return res
}
