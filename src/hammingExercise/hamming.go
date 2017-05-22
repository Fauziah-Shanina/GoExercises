package hammingExercise

import (
	"strings"
)

func HammingDifference(s1, s2 string) (numb int) {
	numb = 0
	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)
	for i := 0; i < len([]rune(s1)); i++ {
		if s1[i] != s2[i] {
			numb++
		}

	}
	return
}
