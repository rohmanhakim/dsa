package lc1268

import (
	"sort"
	"strings"
)

func suggestedProductsBinarySearch(products []string, searchWord string) [][]string {
	sort.Strings(products)
	result := [][]string{}
	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		// Find first product >= prefix
		idx := sort.SearchStrings(products, prefix)
		suggestions := []string{}
		for j := idx; j < len(products) && j < idx+3; j++ {
			if strings.HasPrefix(products[j], prefix) {
				suggestions = append(suggestions, products[j])
			}
		}
		result = append(result, suggestions)
	}
	return result
}
