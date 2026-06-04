package easy

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func evenlyDivideString(s string, num int) string {
	length := len(s)

	if length%num != 0 || num > length {
		return ""
	}

	if num == 1 {
		return s
	}

	newLength := length / num

	i := 0
	p := 0
	substring := ""

	for i < num {
		test := string(s[p : p+newLength])
		if test == substring && i == num-1 {
			return substring
		}
		if substring != "" && test != substring {
			return ""
		}
		substring = test
		p += newLength
		i++
	}
	return ""
}

func canDivide(s string, t string) bool {
	length := len(t)
	if len(s) < length || len(s)%length != 0 {
		return false
	}

	i := 0
	for i < len(s) {
		test := s[i : i+length]
		if test != t {
			return false
		}
		i += length
	}

	return true
}

func gcdOfStrings(str1 string, str2 string) string {

	long := str1
	short := str2

	if len(str2) > len(str1) {
		long = str2
		short = str1
	}

	gcd := ""
	for i := len(short); i > 0; i-- {
		test := evenlyDivideString(short, i)
		if test == "" {
			continue
		}
		if canDivide(long, test) {
			gcd = test
		}
	}

	return gcd
}

func gcdOfStringsClaudeSonnet(str1 string, str2 string) string {
	// If a common divisor exists, concatenation must be commutative
	if str1+str2 != str2+str1 {
		return ""
	}

	// The GCD string length is gcd of the two lengths
	gcdLen := lib.Gcd(len(str1), len(str2))
	return str1[:gcdLen]
}

func GreatestCommonDivisorOfSTrings() {
	// str1 := "ABCABC"
	// str2 := "ABC" // ABC
	// str1 := "LEET"
	// str2 := "CODE" //
	// str1 := "AAAAAB"
	// str2 := "AAA" //
	// str1 := "ABABABAB"
	// str2 := "ABABAB" // "AB"
	// str1 := "ABABABAB"
	// str2 := "ABAB"

	str1 := "AAAAAAAAAAAAAAAAAAAAAAAA"
	str2 := "AAAAAAAAAAAAAAAAAA" //

	// str1 := "AAAAAAAAA"
	// str2 := "AAACCC" //
	// str1 := "ABABAB"
	// str2 := "ABAB" // AB
	out := gcdOfStringsClaudeSonnet(str1, str2)
	fmt.Println(out)

	// fmt.Println(evenlyDivideString("ABAB", 2))
}
