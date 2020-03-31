package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/OakAnderson/ChatBot/util/rand"
	_ "github.com/mattn/go-sqlite3" // Driver para sqlite3
)

// Message guarda valores como o id no banco de dados, a mensagem e
// uma resposta
type Message struct {
	id      int64
	Message string
	Answer  string
}

func connectDatabase(dbName string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", dbName)

	if err != nil {
		return nil, err
	}

	return
}

// QueryMessage recupera uma coluna no banco de dados correspondente a
// expressão recebida
func QueryMessage(db *sql.DB, query string) (mss Message, answers []string, err error) {
	row := db.QueryRow(query)
	answers = make([]string, 5)

	err = row.Scan(&mss.id, &mss.Message, &answers[0], &answers[1], &answers[2],
		&answers[3], &answers[4])

	return
}

// RecoverAnswer retorna uma mensagem encontrada no banco de dados
// para a mensagem recebida
func RecoverAnswer(message string) (mss Message, err error) {
	db, err := connectDatabase("database/messages.db")
	if err != nil {
		return
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM messages WHERE message == \"%s\";", message)

	mss, answers, err := QueryMessage(db, query)

	mss.Answer = rand.SelectString(answers, func(v string) bool {
		return v == "None"
	})

	if mss.Answer == "" {
		return mss, errors.New("nenhuma resposta possível")
	}

	return mss, nil
}
