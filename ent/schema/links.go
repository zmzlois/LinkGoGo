package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Links holds the schema definition for the Links entity.
type Links struct {
	ent.Schema
}

// Fields of the Links.
func (Links) Fields() []ent.Field {
	links := []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
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
			Field("user_id").
			Ref("users_links").
			Required().
			Unique(),
	}
}
