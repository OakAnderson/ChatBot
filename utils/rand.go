package utils

import "math/rand"

func RandAnswer(answers []string) string {
	valids := make([]string, 0, 5)
	for _, v := range answers {
		if v != "None" {
			valids = append(valids, v)
		}
	}

	indx := rand.Intn(len(valids))
	return valids[indx]
}
