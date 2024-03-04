package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	session := []ent.Field{
		field.String("session_token").Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.Time("expires_at"),
	}
	return append(session, BaseTable()...)
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Users.Type).
			Field("user_id").
			Ref("users_sessions").
			Required().
			Unique(),
	}
}
