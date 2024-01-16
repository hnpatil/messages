package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Conversation struct {
	ent.Schema
}

func (Conversation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").
			Default(time.Now()),
		field.String("preview"),
		field.String("conversation_id").
			Unique().
			NotEmpty().
			Immutable(),
	}
}

func (Conversation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("updated_at"),
	}
}

func (Conversation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
