// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LinksColumns holds the columns for the "links" table.
	LinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "url", Type: field.TypeString, Size: 255},
		{Name: "title", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "image", Type: field.TypeString, Size: 255},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// LinksTable holds the schema information for the "links" table.
	LinksTable = &schema.Table{
		Name:       "links",
		Columns:    LinksColumns,
		PrimaryKey: []*schema.Column{LinksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "links_users_users_links",
				Columns:    []*schema.Column{LinksColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "external_id", Type: field.TypeString, Size: 255},
		{Name: "username", Type: field.TypeString, Size: 255},
		{Name: "global_name", Type: field.TypeString, Size: 255},
		{Name: "slug", Type: field.TypeString, Size: 255},
		{Name: "first_name", Type: field.TypeString, Size: 255},
		{Name: "last_name", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Size: 255},
		{Name: "avatar", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Size: 255},
		{Name: "access_token", Type: field.TypeString, Size: 255},
		{Name: "refresh_token", Type: field.TypeString, Size: 255},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LinksTable,
		UsersTable,
	}
)

func init() {
	LinksTable.ForeignKeys[0].RefTable = UsersTable
}
