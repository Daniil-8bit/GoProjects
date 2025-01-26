package sentences

import "strings"

func JoinSomeWords(words []string) string {
	if len(words) == 0 {
		return ""
	} else if len(words) == 1 {
		return words[0]
	} else if len(words) == 2 {
		return words[0] + " and " + words[1]
	} else {
		fewWords := strings.Join(words[:len(words)-1], ", ")
		result := fewWords + ", and " + words[len(words)-1]
		return result
	}
}
