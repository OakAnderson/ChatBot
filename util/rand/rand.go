package rand

import "math/rand"

// SelectString seleciona de forma aleatória um valor do slice
// recebe uma função que descarta valores indesejados
func SelectString(str []string, discart func(v string) bool) string {
	valids := make([]string, 0, len(str))
	for _, v := range str {
		if !discart(v) {
			valids = append(valids, v)
		}
	}

	indx := rand.Intn(len(valids))
	return valids[indx]
}
