package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/OakAnderson/ChatBot/utils"
	_ "github.com/mattn/go-sqlite3"
)

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

func QueryMessage(db *sql.DB, query string) (mss Message, answers []string, err error) {
	row := db.QueryRow(query)
	answers = make([]string, 5)

	err = row.Scan(&mss.id, &mss.Message, &answers[0], &answers[1], &answers[2],
		&answers[3], &answers[4])

	return
}

func RecoverAnswer(message string) (mss Message, err error) {
	db, err := connectDatabase("database/messages.db")
	if err != nil {
		return
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM messages WHERE message == \"%s\";", message)

	mss, answers, err := QueryMessage(db, query)

	mss.Answer = utils.RandAnswer(answers)

	if mss.Answer == "" {
		return mss, errors.New("Nenhuma resposta possível.")
	}

	return mss, nil
}
