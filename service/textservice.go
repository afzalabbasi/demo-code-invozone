package service

import (
	"fmt"
	"github.com/afzalabbasi/demo-code-invozone/textapiroute"
	"regexp"
	"sort"
	"strings"
)

func WordCount(text string) ([]textapiroute.WordCounts, error) {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err

	}
	text = re.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)
	words := strings.Fields(text)

	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	// create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]textapiroute.WordCounts, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, textapiroute.WordCounts{Word: key, Count: val})
	}

	// sort wordCount slice by decreasing count number
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	if len(wordCounts) > 10 {
		return wordCounts[:10], nil
	}

	return wordCounts, nil

}
