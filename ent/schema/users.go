package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	users := []ent.Field{
		field.String("external_id").MaxLen(255),
		field.String("username").MaxLen(255),
		field.String("slug").MaxLen(255),
		field.String("first_name").MaxLen(255),
		field.String("last_name").MaxLen(255),
		field.String("email").MaxLen(255),
		field.String("avatar").MaxLen(255),
		field.String("description").MaxLen(255),
		field.String("access_token").MaxLen(255),
		field.String("refresh_token").MaxLen(255),
	}
	return append(users, BaseTable()...)
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users_links", Links.Type),
	}
}
