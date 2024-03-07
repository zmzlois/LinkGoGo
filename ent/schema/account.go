package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account. user sets their own user data here, if we have data from this table we won't be looking for users table because those coming from discord.
// This account table is supposed to override user table
func (Account) Fields() []ent.Field {
	account := []ent.Field{
		field.UUID("external_id", uuid.UUID{}),
		field.String("username").Unique(),
		field.String("first_name").MaxLen(255).Optional(),
		field.String("last_name").MaxLen(255).Optional(),
		field.String("email").MaxLen(255).Optional(),
		field.String("avatar").MaxLen(255).Optional(),
	}
	return append(account, BaseTable()...)
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).
			Field("external_id").
			Ref("users_accounts").
			Required().Unique(),
	}
}
