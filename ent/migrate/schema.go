// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "external_id", Type: field.TypeUUID},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_users_accounts",
				Columns:    []*schema.Column{AccountsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LinksColumns holds the columns for the "links" table.
	LinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "url", Type: field.TypeString, Size: 255},
		{Name: "title", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "image", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "order", Type: field.TypeInt, Default: 0},
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
				Columns:    []*schema.Column{LinksColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "session_token", Type: field.TypeString, Unique: true},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "external_id", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Size: 255},
		{Name: "global_name", Type: field.TypeString, Size: 255},
		{Name: "slug", Type: field.TypeString, Size: 255},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "access_token", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "refresh_token", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "scope", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "expires_in", Type: field.TypeFloat64, Nullable: true},
		{Name: "session_state", Type: field.TypeString, Nullable: true},
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
		AccountsTable,
		LinksTable,
		SessionsTable,
		UsersTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	LinksTable.ForeignKeys[0].RefTable = UsersTable
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
}
