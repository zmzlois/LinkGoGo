package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

func BaseTable() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			MaxLen(255),
		field.Bool("deleted").
			Default(false),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}),
	}
}
