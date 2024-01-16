package repository

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hnpatil/messages/entity"
	"github.com/hnpatil/messages/entity/conversation"
	"github.com/hnpatil/messages/entity/message"
	"github.com/hnpatil/messages/entity/user"
)

type messagesImpl struct {
	db *entity.Client
}

func NewMessages(db *entity.Client) Messages {
	return &messagesImpl{
		db: db,
	}
}

func (m *messagesImpl) CreateMessage(ctx context.Context, message *entity.Message, convID string) (*entity.Message, error) {
	tx, err := m.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	conv, err := tx.Conversation.Create().
		SetConversationID(convID).
		SetUpdatedAt(now).
		SetPreview(message.Text).
		OnConflict(sql.ConflictColumns(conversation.FieldConversationID)).
		UpdateNewValues().
		ID(ctx)

	if err != nil {
		return nil, err
	}

	msg, err := tx.Message.Create().
		SetText(message.Text).
		SetSender(message.Edges.Sender).
		SetRecipient(message.Edges.Recipient).
		SetCreatedAt(now).
		SetConversationID(conv).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return msg, err
}

func (m *messagesImpl) ListConversations(ctx context.Context, usr *entity.User) ([]*entity.Conversation, error) {
	return m.db.Conversation.Query().
		Where(
			conversation.HasMessagesWith(
				message.Or(
					message.HasRecipientWith(user.ID(usr.ID)),
					message.HasSenderWith(user.ID(usr.ID)),
				),
			),
		).
		Order(conversation.ByUpdatedAt(
			sql.OrderDesc(),
		)).WithMessages(func(mq *entity.MessageQuery) { mq.Order(message.ByCreatedAt(sql.OrderDesc())).Limit(1) }).
		All(ctx)
}

func (m *messagesImpl) ListMessages(ctx context.Context, usr *entity.User, convID string) ([]*entity.Message, error) {
	return m.db.Message.Query().
		Where(
			message.HasConversationWith(conversation.ConversationID(convID)),
			message.Or(
				message.HasRecipientWith(user.ID(usr.ID)),
				message.HasSenderWith(user.ID(usr.ID)),
			),
		).
		Order(message.ByCreatedAt(
			sql.OrderDesc(),
		)).
		All(ctx)
}
