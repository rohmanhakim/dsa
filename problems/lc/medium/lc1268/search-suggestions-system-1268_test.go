package lc1268

import (
	"reflect"
	"testing"
)

func TestSearchSuggestionsSystem(t *testing.T) {
	tests := []struct {
		name       string
		products   []string
		searchWord string
		want       [][]string
	}{
		{
			name:       "example 1",
			products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			searchWord: "mouse",
			want:       [][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}},
		},
		{
			name:       "example 2",
			products:   []string{"havana"},
			searchWord: "havana",
			want:       [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
		{
			name:       "example 3",
			products:   []string{"havana"},
			searchWord: "tatiana",
			want:       [][]string{{}, {}, {}, {}, {}, {}, {}},
		},
		{
			name:       "example 4",
			products:   []string{"ilpsokmmniywxgbeqrpsaqeqn", "ilpxatcfeexjqxghbengdcvsljajqaxidzitqjfjpovxmijvofntfelqidcbekzecqodiralswkjqckrpz", "ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr", "ilpxatcgvn", "ilpxatcgvtqdvgovjgabbmkdicgymudbikvhihpxfmpanjyp", "ilpxatcsfwrpusskcsmzunvejmfymcvs", "ilpxxishsvmt", "ilpxywtmdnlil", "irjbjafrvdcpqanwt"},
			searchWord: "ilpxatcgvfblxfvbnoxgmmpvimowo",
			want:       [][]string{{"ilpsokmmniywxgbeqrpsaqeqn", "ilpxatcfeexjqxghbengdcvsljajqaxidzitqjfjpovxmijvofntfelqidcbekzecqodiralswkjqckrpz", "ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpsokmmniywxgbeqrpsaqeqn", "ilpxatcfeexjqxghbengdcvsljajqaxidzitqjfjpovxmijvofntfelqidcbekzecqodiralswkjqckrpz", "ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpsokmmniywxgbeqrpsaqeqn", "ilpxatcfeexjqxghbengdcvsljajqaxidzitqjfjpovxmijvofntfelqidcbekzecqodiralswkjqckrpz", "ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpxatcfeexjqxghbengdcvsljajqaxidzitqjfjpovxmijvofntfelqidcbekzecqodiralswkjqckrpz", "ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr", "ilpxatcgvn"}, {"ilpxatcfeexjqxghbengdcvsljajqaxidzitqjfjpovxmijvofntfelqidcbekzecqodiralswkjqckrpz", "ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr", "ilpxatcgvn"}, {"ilpxatcfeexjqxghbengdcvsljajqaxidzitqjfjpovxmijvofntfelqidcbekzecqodiralswkjqckrpz", "ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr", "ilpxatcgvn"}, {"ilpxatcfeexjqxghbengdcvsljajqaxidzitqjfjpovxmijvofntfelqidcbekzecqodiralswkjqckrpz", "ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr", "ilpxatcgvn"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr", "ilpxatcgvn", "ilpxatcgvtqdvgovjgabbmkdicgymudbikvhihpxfmpanjyp"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr", "ilpxatcgvn", "ilpxatcgvtqdvgovjgabbmkdicgymudbikvhihpxfmpanjyp"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {"ilpxatcgvfblxfvbnjyryxxfayspcrqkvyopfzspzmmhhalwjvfhsgybgkzctlqtr"}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := suggestedProducts(tt.products, tt.searchWord)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}

func TestSearchSuggestionsBinarySearch(t *testing.T) {
	tests := []struct {
		name       string
		products   []string
		searchWord string
		want       [][]string
	}{
		{
			name:       "example 1",
			products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			searchWord: "mouse",
			want:       [][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}},
		},
		{
			name:       "example 2",
			products:   []string{"havana"},
			searchWord: "havana",
			want:       [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
		{
			name:       "example 3",
			products:   []string{"havana"},
			searchWord: "tatiana",
			want:       [][]string{{}, {}, {}, {}, {}, {}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := suggestedProducts(tt.products, tt.searchWord)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
