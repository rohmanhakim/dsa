package medium

import "fmt"

func predictPartyVictory(senate string) string {
	votes := []rune{}

	d := 0
	r := 0
	for d < len(senate) || r < len(senate) {
		sfront := rune(senate[0])
		var qfront rune
		banned := false
		if len(votes) > 0 {
			qfront = votes[0]
			if qfront != sfront {
				votes = votes[1:]
				banned = true
			}
		}

		senate = senate[1:]
		if !banned {
			senate += string(sfront)
			votes = append(votes, sfront)
		}

		if len(votes) == len(senate) {
			if sfront == 'D' {
				return "Dire"
			} else {
				return "Radiant"
			}
		}
	}

	return ""
}

func Dota2Senate() {
	// senate := "DDRRR" // Dire
	// senate := "DR" // Dire
	// senate := "RD" // Radiant
	senate := "RDD" // Dire

	out := predictPartyVictory(senate)
	fmt.Println(out)
}
