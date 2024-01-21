package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
		field.String("text"),
	}
}

func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipient", User.Type).
			Ref("received_message").
			Unique().
			Required().
			Immutable(),

		edge.From("sender", User.Type).
			Ref("sent_message").
			Unique().
			Required().
			Immutable(),

		edge.From("conversation", Conversation.Type).
			Ref("messages").
			Unique(),
	}
}
