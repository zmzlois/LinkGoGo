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
		field.String("external_id").Unique(),
		field.String("username").MaxLen(255),
		field.String("global_name").MaxLen(255),
		field.String("slug").MaxLen(255),
		field.String("first_name").MaxLen(255).Optional(),
		field.String("last_name").MaxLen(255).Optional(),
		field.String("email").MaxLen(255).Optional(),
		field.String("avatar").MaxLen(255).Optional(),
		field.String("description").MaxLen(255).Optional(),
		field.String("access_token").MaxLen(255).Optional(),
		field.String("refresh_token").MaxLen(255).Optional(),
		field.String("scope").MaxLen(255).Optional(),
		field.Float("expires_in").Optional(),
		field.String("session_state").Optional(),
	}
	return append(users, BaseTable()...)
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users_links", Links.Type),
		edge.To("users_sessions", Session.Type),
	}
}
