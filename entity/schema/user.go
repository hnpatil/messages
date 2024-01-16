package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name"),
		field.String("last_name"),
		field.String("email").
			Unique(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("received_message", Message.Type),
		edge.To("sent_message", Message.Type),
	}
}
