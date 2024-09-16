package listener

type WebhookListener func(message *WebhookMessage) error

type WebhookMessage struct {
	IsStatusReply  bool                 `json:"isStatusReply"`
	ChatLid        string               `json:"chatLid"`
	ConnectedPhone string               `json:"connectedPhone"`
	WaitingMessage bool                 `json:"waitingMessage"`
	IsEdit         bool                 `json:"isEdit"`
	IsGroup        bool                 `json:"isGroup"`
	IsNewsletter   bool                 `json:"isNewsletter"`
	InstanceId     string               `json:"instanceId"`
	MessageId      string               `json:"messageId"`
	Phone          string               `json:"phone"`
	FromMe         bool                 `json:"fromMe"`
	Momment        int64                `json:"momment"`
	Status         string               `json:"status"`
	ChatName       string               `json:"chatName"`
	SenderPhoto    interface{}          `json:"senderPhoto"`
	SenderName     string               `json:"senderName"`
	Photo          string               `json:"photo"`
	Broadcast      bool                 `json:"broadcast"`
	ParticipantLid interface{}          `json:"participantLid"`
	Forwarded      bool                 `json:"forwarded"`
	Type           string               `json:"type"`
	FromApi        bool                 `json:"fromApi"`
	Text           *WebhookMessageText  `json:"text"`
	Audio          *WebhookMessageAudio `json:"audio"`
}

type WebhookMessageAudio struct {
	Ptt      bool   `json:"ptt"`
	Seconds  int    `json:"seconds"`
	AudioUrl string `json:"audioUrl"`
	MimeType string `json:"mimeType"`
	ViewOnce bool   `json:"viewOnce"`
}

type WebhookMessageText struct {
	Message string `json:"message"`
}
