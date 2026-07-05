package lc383

/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 10^5
ransomNote and magazine consist of lowercase English letters.
*/
import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	freq := make(map[rune]int)
	for _, ch := range magazine {
		freq[ch]++
	}

	for _, ch := range ransomNote {
		if freq[ch] == 0 {
			return false
		}
		freq[ch]--
	}
	return true
}

func LC383() {
	ransomNote := "a"
	magazine := "b"
	fmt.Println(canConstruct(ransomNote, magazine))
}
