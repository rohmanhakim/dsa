package lc13

func romanToInt(s string) int {

	symbols := make(map[rune]int)
	symbols['I'] = 1
	symbols['V'] = 5
	symbols['X'] = 10
	symbols['L'] = 50
	symbols['C'] = 100
	symbols['D'] = 500
	symbols['M'] = 1000

	needSubstract := make(map[rune]rune)
	needSubstract['V'] = 'I'
	needSubstract['X'] = 'I'
	needSubstract['L'] = 'X'
	needSubstract['C'] = 'X'
	needSubstract['D'] = 'C'
	needSubstract['M'] = 'C'

	res := 0
	var prev rune
	for i, ch := range s {
		if i == 0 {
			prev = ch
			res += symbols[ch]
		} else if val, ok := needSubstract[ch]; ok && prev == val {
			res -= symbols[prev]
			res += symbols[ch] - symbols[prev]
		} else {
			res += symbols[ch]
		}
		prev = ch
	}
	return res
}
