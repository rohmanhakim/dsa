package medium

import "fmt"

func jumpBackTrack(arr []int, i int, path *[]int, found *bool, visited []int) {
	visited[i] = 1
	pathCopy := *path
	if arr[i] == 0 {
		*found = true
		return
	} else if len(pathCopy) > 2 && i == pathCopy[0] {
		*path = pathCopy[:len(pathCopy)-1]
		return
	}

	left := i - arr[i]
	right := i + arr[i]

	choice := []int{}
	if left >= 0 && visited[left] == 0 {
		choice = append(choice, left)
	}
	if right <= len(arr)-1 && visited[right] == 0 {
		choice = append(choice, right)
	}

	for _, c := range choice {
		*path = append(*path, c)
		jumpBackTrack(arr, c, path, found, visited)
		if len(*path) == 0 {
			return
		}
		tmp := *path
		*path = tmp[:len(tmp)-1]
	}
}

func canReach(arr []int, start int) bool {
	found := false
	path := []int{start}
	visited := make([]int, len(arr))
	jumpBackTrack(arr, start, &path, &found, visited)
	return found
}

func canReachClaudeSonnet(arr []int, start int) bool {
	n := len(arr)
	visited := make([]bool, n)
	queue := []int{start}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]

		if arr[i] == 0 {
			return true
		}
		if visited[i] {
			continue
		}
		visited[i] = true

		if left := i - arr[i]; left >= 0 && !visited[left] {
			queue = append(queue, left)
		}
		if right := i + arr[i]; right < n && !visited[right] {
			queue = append(queue, right)
		}
	}
	return false
}

func JumpGame3() {
	arr := []int{4, 2, 3, 0, 3, 1, 2} // true
	start := 3
	// arr := []int{4, 2, 3, 0, 3, 1, 2} // true
	// start := 0
	// arr := []int{3, 0, 2, 1, 2} // false
	// start := 2
	// arr := []int{4, 4, 1, 3, 0, 3} // true
	// start := 2
	// arr := []int{3, 0, 2, 1, 2} // false
	// start := 2
	// arr := []int{4, 2, 3, 0, 3, 1, 2} // true
	// start := 0
	// arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 0}
	// start := 1
	out := canReach(arr, start)
	fmt.Println(out)
}
