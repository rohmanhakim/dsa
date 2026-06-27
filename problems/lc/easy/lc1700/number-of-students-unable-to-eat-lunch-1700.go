package lc1700

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	unable := 0
	for len(students) > 0 {
		if len(sandwiches) > 0 {
			if sandwiches[0] == students[0] {
				unable = 0
				sandwiches = sandwiches[1:]
				students = students[1:]
			} else {
				unable++
				front := students[0]
				students = students[1:]
				students = append(students, front)
			}
			if unable == len(students) {
				break
			}
		} else {
			break
		}
	}

	return len(students)
}

func NumberOfStudentsUnableTOEatLunch() {
	// students := []int{1, 1, 0, 0}
	// sandwiches := []int{0, 1, 0, 1}
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	// 11001 00011 1
	// 10011 00011 2
	// 00111 00011 0
	// 0111 0011 0
	// 111 011 1
	// 111 011 2

	out := countStudents(students, sandwiches)
	fmt.Println(out)
}
