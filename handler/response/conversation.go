package response

import (
	"time"

	"github.com/hnpatil/messages/entity"
)

type Conversation struct {
	ID          string    `json:"id,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Preview     string    `json:"preview,omitempty"`
	LastMessage *Message  `json:"last_message,omitempty"`
}

func ToListConversationsResponse(input []*entity.Conversation) []*Conversation {
	res := make([]*Conversation, len(input))

	for i := 0; i < len(input); i++ {
		res[i] = ToConversationResponse(input[i])
	}

	return res
}

func ToConversationResponse(input *entity.Conversation) *Conversation {
	c := &Conversation{
		ID:        input.ConversationID,
		UpdatedAt: input.UpdatedAt,
		Preview:   input.Preview,
	}

	if len(input.Edges.Messages) > 0 {
		c.LastMessage = ToMessageResponse(input.Edges.Messages[0])
	}

	return c
}
