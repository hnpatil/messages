package usecase

import (
	"fmt"
	"sort"

	"github.com/google/uuid"
	"github.com/hnpatil/messages/entity"
	"github.com/hnpatil/messages/repository"
)

type messagesImpl struct {
	messages repository.Messages
	users    repository.Users
}

func NewMessages(messages repository.Messages, users repository.Users) Messages {
	return &messagesImpl{
		messages: messages,
		users:    users,
	}
}

func (m *messagesImpl) SendMessage(ctx Context, recipientID string, text string) (*entity.Message, error) {
	senderID := ctx.GetUserID()
	sender, err := m.users.GetUser(ctx.GetContext(), senderID)
	if err != nil {
		return nil, err
	}

	recipient, err := m.users.GetUser(ctx.GetContext(), recipientID)
	if err != nil {
		return nil, err
	}

	convID := m.getConversationID(sender, recipient)

	message := &entity.Message{
		Text: text,
		Edges: entity.MessageEdges{
			Sender:    sender,
			Recipient: recipient,
		},
	}

	return m.messages.CreateMessage(ctx.GetContext(), message, convID)
}

func (m *messagesImpl) ListConversations(ctx Context) ([]*entity.Conversation, error) {
	usrID := ctx.GetUserID()
	usr, err := m.users.GetUser(ctx.GetContext(), usrID)
	if err != nil {
		return nil, err
	}

	return m.messages.ListConversations(ctx.GetContext(), usr)
}

func (m *messagesImpl) ListMessages(ctx Context, forConversation string) ([]*entity.Message, error) {
	usrID := ctx.GetUserID()
	usr, err := m.users.GetUser(ctx.GetContext(), usrID)
	if err != nil {
		return nil, err
	}

	return m.messages.ListMessages(ctx.GetContext(), usr, forConversation)
}

func (m *messagesImpl) getConversationID(usr ...*entity.User) string {
	sort.Slice(usr, func(i, j int) bool {
		return usr[i].ID < usr[j].ID
	})

	convID := ""

	for _, it := range usr {
		convID = fmt.Sprintf("%s%s", convID, it.Email)
	}

	id := uuid.NewSHA1(uuid.UUID{}, []byte(convID))

	return id.String()
}
