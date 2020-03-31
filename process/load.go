package process

import (
	s "strings"

	db "github.com/OakAnderson/ChatBot/database"
	"github.com/OakAnderson/ChatBot/util/str"
)

// LoadMessage processa a atual informação e retorna uma resposta apropriada
func LoadMessage(message string) string {
	mss, err := db.RecoverAnswer(s.ToLower(message))
	if err != nil {
		return ""
	}

	return str.Capitalize(mss.Answer)
}

// LoadCommand processa a mensagem como um comando predefinido
func LoadCommand(message string) string {
	if s.HasPrefix(message, "/dado") {
		return ComandoDado(message)
	}

	return "Comando não encontrado."
}

// LoadReply verifica o tipo de mensagem recebida e chama a função referente
func LoadReply(message string) string {
	if s.HasPrefix(message, "/") {
		return LoadCommand(message)
	}

	return LoadMessage(message)
}
