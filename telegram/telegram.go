package telegram

const TelegramLink = "https://api.telegram.org"
const BotToken = "1037487028:AAHUG8-W3ufixfrmAejOYjls0P_V8sWoBC4"

type WebhookReqBody struct {
	Message struct {
		Message_id int32 `json:"message_id"`
		From       struct {
			ID            int64  `json:"id"`
			Is_bot        bool   `json:"is_bot"`
			First_name    string `json:"first_name"`
			Last_name     string `json:"last_name"`
			Username      string `json:"username"`
			Language_code string `json:"pt-br"`
		} `json:"from"`
		Chat struct {
			ID         int64  `json:"id"`
			First_name string `json:"first_name"`
			Last_name  string `json:"last_name"`
			Username   string `json:"username"`
			Type       string `json:"private"`
		} `json:"chat"`
		Text string `json:"text"`
		Date int64  `json:"date"`
	} `json:"message"`
}

type SendMensageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
