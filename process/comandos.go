package process

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ComandoDado(message string) string {
	args := strings.Split(message, " ")
	max, num := 6, 1
	var err error

	if len(args) >= 3 {
		num, err = strconv.Atoi(args[2])
		if err != nil {
			return "Argumentos inválidos"
		}
	}
	if len(args) >= 2 {
		max, err = strconv.Atoi(args[1])
		if err != nil {
			return "Argumentos inválidos"
		}
	}

	dados := make([]string, num)
	for i := 0; i < num; i++ {
		dados[i] = strconv.Itoa(rand.Intn(max) + 1)
	}

	return strings.Join(dados, " ")
}
