package lc735

import "fmt"

func abs(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}
func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, a := range asteroids {
		if len(stack) == 0 {
			// if empty any item can be pushed
			stack = append(stack, a)
			continue
		}
		if len(stack) > 0 { // if there are items in the stack
			// top := stack[len(stack)-1]

			// - negatives can only be pushed if no higher/equal positives in the stack
			// - positives can be pushed anytime
			// - destroying items (in the stack/incoming) can only be triggered if new item is negative
			//   and there is positive in the stack

			if a > 0 {
				stack = append(stack, a)
			}
			if a < 0 {
				push := true
				for len(stack) > 0 &&
					stack[len(stack)-1] > 0 &&
					abs(a) >= stack[len(stack)-1] {
					top := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if abs(a) == top {
						push = false
						break
					}
				}
				if len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] > abs(a) {
					push = false
				}
				if push {
					stack = append(stack, a)
				}
			}
		}
	}
	return stack
}

func asteroidCollisionClaudeSonnet(asteroids []int) []int {
	stack := []int{}

	for _, a := range asteroids {
		alive := true

		for alive && a < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top <= -a {
				stack = stack[:len(stack)-1] // pop the top
				if top == -a {
					alive = false // mutual destruction
				}
			} else {
				alive = false // incoming asteroid is smaller, it dies
			}
		}

		if alive {
			stack = append(stack, a)
		}
	}

	return stack
}

func AsteroidCollision() {
	asteroids := []int{-2, -2, 1, -2} // -2, -2, -2
	// asteroids := []int{10, 2, -5} // 10
	// asteroids := []int{5, 10, -5} // 5, 10
	// asteroids := []int{8, -8} // []
	// asteroids := []int{3, 5, -6, 2, -1, 4} // -6, 2, 4
	// asteroids := []int{-2, -1, 1, 2} // -2, -1, 1, 2
	// asteroids := []int{-2, 1, -1, -2} // -2, -2
	// asteroids := []int{-2, 1, 1, -1} // -2, 1
	// asteroids := []int{-2, 2, 1, -2} // -2
	out := asteroidCollision(asteroids)
	fmt.Println(out)
}
