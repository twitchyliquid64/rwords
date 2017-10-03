package rwords

import (
	"sort"
	"strings"
)

func buildMap(input string) map[byte][]wordWeight {
	input = strings.NewReplacer(".", " ", ",", " ", "\n", " ", "\t", " ", "'", "", "\"", "").Replace(strings.ToLower(input))
	spl := strings.Split(input, " ")
	incidenceMap := map[byte][]wordWeight{}
	for _, word := range spl {
		if word == "" || word == "\n" {
			continue
		}
		word = strings.TrimSpace(word)

		for i := 0; i < len(word); i++ {
			c := word[i]
			var next string
			if (i + 1) < len(word) {
				next = string(word[i+1])
			}

			found := false
			for i := range incidenceMap[c] {
				if incidenceMap[c][i].word == string(next) {
					incidenceMap[c][i].weight++
					found = true
					break
				}
			}
			if !found {
				incidenceMap[c] = append(incidenceMap[c], wordWeight{next, 1})
			}
		}
	}
	return incidenceMap
}

func sortMap(in map[byte][]wordWeight) map[byte][]wordWeight {
	out := map[byte][]wordWeight{}
	for k, v := range in {
		s := v
		sort.Sort(ByWeight(s))
		out[k] = s
	}
	return out
}
