package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/hnpatil/messages/entity"
)

type Message struct {
	ID        uuid.UUID `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Text      string    `json:"text,omitempty"`
	Sender    *User     `json:"sender"`
	Recipient *User     `json:"recipient"`
}

func ToListMessagesResponse(input []*entity.Message) []*Message {
	res := make([]*Message, len(input))

	for i := 0; i < len(input); i++ {
		res[i] = ToMessageResponse(input[i])
	}

	return res
}

func ToMessageResponse(input *entity.Message) *Message {
	m := &Message{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		Text:      input.Text,
	}

	if sender := input.Edges.Sender; sender != nil {
		m.Sender = ToUserResponse(sender)
	}

	if recipient := input.Edges.Recipient; recipient != nil {
		m.Recipient = ToUserResponse(recipient)
	}

	return m
}
