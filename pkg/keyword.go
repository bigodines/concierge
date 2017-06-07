package keyword

import (
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
