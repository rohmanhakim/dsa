package medium

import "fmt"

func visitAllRooms(rooms [][]int, room int, visited []bool) {

	visited[room] = true

	for _, neighbor := range rooms[room] {
		if !visited[neighbor] {
			visitAllRooms(rooms, neighbor, visited)
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visitAllRooms(rooms, 0, visited)

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func canVisitAllRoomsIterative(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	stack := []int{0}

	for len(stack) > 0 {
		room := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[room] {
			continue
		}
		visited[room] = true

		for _, key := range rooms[room] {
			if !visited[key] {
				stack = append(stack, key)
			}
		}
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func KeysAndRooms() {
	rooms := [][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}
	out := canVisitAllRooms(rooms)
	fmt.Println(out)
}
