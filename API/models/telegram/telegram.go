package telegram

const (
	// TelegramLink é o link para a api do telegram
	TelegramLink = "https://api.telegram.org"
	// BotToken é o token de acesso da API
	BotToken = "1037487028:AAHUG8-W3ufixfrmAejOYjls0P_V8sWoBC4"
)

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
