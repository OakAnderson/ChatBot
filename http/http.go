package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/OakAnderson/ChatBot/process"
	tm "github.com/OakAnderson/ChatBot/telegram"
)

const PORT = ":3000"

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &tm.WebhookReqBody{}

	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println("Could not decode request body:\n\t>>>", err)
		return
	}

	text := process.LoadReply(body.Message.Text)

	if err := sendAnswer(body.Message.Chat.ID, text); err != nil {
		log.Println(err)
		return
	}

	log.Println("Sent reply")
}

func sendAnswer(ChatID int64, Text string) error {
	reqBody := &tm.SendMensageReqBody{
		ChatID: ChatID,
		Text:   Text,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	url := tm.TelegramLink + "/bot" + tm.BotToken + "/sendMessage"

	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Unespected status " + res.Status)
	}

	return nil
}
