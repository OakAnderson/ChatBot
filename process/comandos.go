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

// ComandoDado gera um ou mais numeros aleatorios de acordo com os parametros
func ComandoDado(message string) string {
	args := strings.Split(message, " ")
	max, num := 6, 1
	var err error

	if len(args) >= 3 {
		num, err = strconv.Atoi(args[2])
		if err != nil {
			return "uso: /dado valorMaximo numeroDeDados"
		}
	}
	if len(args) >= 2 {
		max, err = strconv.Atoi(args[1])
		if err != nil {
			return "uso: /dado valorMaximo numeroDeDados"
		}
	}

	dados := make([]string, num)
	for i := 0; i < num; i++ {
		dados[i] = strconv.Itoa(rand.Intn(max) + 1)
	}

	return strings.Join(dados, " ")
}
