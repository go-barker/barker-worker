package telegramsender

import "github.com/corporateanon/barker/pkg/types"

type sendMessagePayload struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func createSendMessagePayload(user *types.User, campaign *types.Campaign) *sendMessagePayload {
	return &sendMessagePayload{
		ChatID: user.TelegramID,
		Text:   campaign.Message,
	}
}
