package telegram

import "io/ioutil"

// TelegramLink é o link para a api do telegram
const TelegramLink = "https://api.telegram.org"

// WebhookReqBody é uma estrura que guarda valores de uma mensagem recebida
type WebhookReqBody struct {
	Message struct {
		MessageID int32 `json:"message_id"`
		From      struct {
			ID           int64  `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"pt-br"`
		} `json:"from"`
		Chat struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"private"`
		} `json:"chat"`
		Text string `json:"text"`
		Date int64  `json:"date"`
	} `json:"message"`
}

// SendMensageReqBody guarda valores json para enviar uma mensagem
type SendMensageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// Token retorna o token do bot
func Token() (token string, err error) {
	tokenFile, err := ioutil.ReadFile("API/models/telegram/token.txt")

	if err != nil {
		return "", err
	}

	return string(tokenFile), nil
}
