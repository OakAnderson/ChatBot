package process

import (
	s "strings"

	db "github.com/OakAnderson/ChatBot/database"
	"github.com/OakAnderson/ChatBot/utils"
)

func LoadMessage(message string) string {
	mss, err := db.RecoverAnswer(s.ToLower(message))
	if err != nil {
		return ""
	}

	return utils.Capitalize(mss.Answer)
}

func LoadCommand(message string) string {
	if s.HasPrefix(message, "/dado") {
		return ComandoDado(message)
	}

	return "Comando não encontrado."
}

func LoadReply(message string) string {
	if s.HasPrefix(message, "/") {
		return LoadCommand(message)
	} else {
		return LoadMessage(message)
	}
}
