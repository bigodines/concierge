package keyword

import (
	"sort"
	"strings"
)

type Matcher struct {
	keywords  []string
	excludes  []string
	sensitive bool
}

// whatever
func New(words string, excludes string, sensitive bool) *Matcher {

	if !sensitive {
		words = strings.ToLower(words)
		excludes = strings.ToLower(excludes)
	}

	wordsArr := strings.Split(words, ",")
	excludesArr := strings.Split(excludes, ",")

	wordsArr = CleanUp(wordsArr)
	excludesArr = CleanUp(excludesArr)

	return &Matcher{
		keywords:  wordsArr,
		excludes:  excludesArr,
		sensitive: sensitive,
	}
}

func (m *Matcher) CheckAll(s string) bool {
	words := strings.Split(s, " ")
	sort.Strings(words)
	matches := 0
	for i := 0; i < len(m.keywords); i++ {
		if indexOf(m.keywords[i], words) >= 0 {
			matches++
		}
	}

	return matches == len(m.keywords)
}

func (m *Matcher) CheckAny(s string) bool {
	return false
}

/**
Removes trailing spaces and fix case in an array of strings
*/
func CleanUp(words []string) []string {
	var ret []string

	for i := 0; i < len(words); i++ {
		str := strings.Trim(words[i], " ")

		if str == "" {
			continue
		}

		ret = append(ret, str)
	}

	return ret
}

/**
Basic implementation of indexOf. Plenty of room for optimization.
*/
func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
