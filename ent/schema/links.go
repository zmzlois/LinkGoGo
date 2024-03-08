package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Links holds the schema definition for the Links entity.
type Links struct {
	ent.Schema
}

type order int

const (
	Order = iota
)

func (Links) Index() []ent.Index {
	return []ent.Index{
		index.Fields("order"),
	}
}

// Fields of the Links.
func (Links) Fields() []ent.Field {
	links := []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.String("url").MaxLen(255),
		field.String("title").MaxLen(255),
		field.Text("description").Optional(),
		field.String("image").MaxLen(255).Optional(),
		field.Int("order").Default(0),
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
