package lc790

import "fmt"

const MOD = 1000000007

func numTilingsTopDown(n int) int {
	// Memoization arrays: -1 means not computed yet
	memo := make([]int, n+1)
	memoP := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
		memoP[i] = -1
	}

	var dp func(n int) int
	var p func(n int) int

	dp = func(n int) int {
		if n < 0 {
			return 0
		}
		if memo[n] != -1 {
			return memo[n]
		}

		var res int
		switch n {
		case 0:
			res = 1 // empty board: only 1 possibilities to complete board: place nothing
		case 1:
			res = 1 // one possibility to create complete board: one vertical domino
		case 2:
			res = 2 // two valid possibilities to create complete board: two verticals and two horizontal dominoes
		default:
			// Recurrence: dp(n) = dp(n-1) + dp(n-2) + p(n-2)
			// 
			// Explanation of p(n-2) term:
			// To create a complete board at n using trominos, we must close a "gap" 
			// created by a previous tromino. If we have an overhang state p(n-2), 
			// that means one cell of column (n-1) is already covered by a tromino 
			// extending from the left. We place a second tromino covering:
			// - The empty cell in column (n-1) 
			// - Both cells in column n
			// This fills the gap and completes the board at column n.
			// Visual:
			//   Column:  n-2   n-1    n
			//           ...   [X]   [?]   <- [X] pre-covered by tromino from n-2
			//           ...   [.]   [?]   <- [.] is the gap, [?] are empty
			//   Place tromino: covers [.] + both [?][?] = complete board dp(n)!
			res = (dp(n-1) + dp(n-2) + p(n-2)) % MOD
		}

		memo[n] = res
		return res
	}

	p = func(n int) int {
		if n < 0 {
			return 0
		}
		if memoP[n] != -1 {
			return memoP[n]
		}

		var res int
		switch n {
		case 0:
			res = 0 // zero posibilities to create overhang board in empty board
		case 1:
			/*
				two possibilities of creating overhang board:
				AA
				A

				A
				AA
			*/

			res = 2
		case 2:
			/*
			   four possibilities ove creating overhang board:
			   AA
			   ABB

			   ABB
			   AA

			   ABB
			   AB

			   AB
			   ABB
			*/

			res = 4
		default:
			// Recurrence: p(n) = 2*dp(n-1) + p(n-1)
			//
			// Explanation of 2*dp(n-1) term:
			// To create an overhang at column (n+1), we start from a complete board 
			// at (n-1) [dp(n-1)] and place a tromino covering both cells of column n 
			// and one cell of column (n+1). 
			// 
			// The factor of 2 comes from two symmetric orientations:
			// 1. Tromino covers top of column (n+1): L-shape pointing up-right
			//    Column n   n+1
			//       AA       A    
			//       AA       .    
			// 2. Tromino covers bottom of column (n+1): L-shape pointing down-right  
			//    Column n   n+1
			//       BB       .    
			//       BB       B    
			//
			// Both create valid p(n) states (overhang at n+1).
			//
			// Explanation of p(n-1) term:
			// If we already have an overhang at column n [p(n-1)], we can place a 
			// horizontal domino covering the empty cell in column n and extending 
			// into column (n+1). This fills column n completely but creates a new 
			// overhang at column (n+1), propagating the gap forward.
			res = (2*dp(n-1) + p(n-1)) % MOD
		}

		memoP[n] = res
		return res
	}

	return dp(n)
}

func numTilings(n int) int {
	out := numTilingsTopDown(n)
	return out
}

func DominoAndTrominoTiling() {
	n := 3
	out := numTilings(n)
	fmt.Println(out)
}
