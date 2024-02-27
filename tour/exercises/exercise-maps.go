package tour

import (
	"strings"
)

func WordCount(s string) map[string]int {
	mapWords := make(map[string]int)
	sliceWords := strings.Fields(s)

	for _, word := range sliceWords {
		value, ok := mapWords[word]

		if ok {
			mapWords[word] += value + 1
		} else {
			mapWords[word] = 1
		}
	}

	return mapWords
}

// func main() {
// 	wc.Test(WordCount)
// }
