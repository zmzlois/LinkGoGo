package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Links holds the schema definition for the Links entity.
type Links struct {
	ent.Schema
}

// Fields of the Links.
func (Links) Fields() []ent.Field {
	links := []ent.Field{
		field.String("user_id").MaxLen(255),
		field.String("url").MaxLen(255),
		field.String("title").MaxLen(255).Unique(),
		field.String("image").MaxLen(255),
	}

	return append(links, BaseTable()...)
}

// Edges of the Links.
func (Links) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Users.Type).
			Ref("users_links").
			Unique(),
	}
}
