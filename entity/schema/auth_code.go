package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AuthCode struct {
	ent.Schema
}

func (AuthCode) Fields() []ent.Field {
	return []ent.Field{
		field.Time("expires_at"),
		field.String("identifier"),
		field.String("code").
			Unique(),
	}
}

func (AuthCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("identifier"),
	}
}
