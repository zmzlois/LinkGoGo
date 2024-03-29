// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/zmzlois/LinkGoGo/ent/users"
)

// Users is the model entity for the Users schema.
type Users struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ExternalID holds the value of the "external_id" field.
	ExternalID string `json:"external_id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// GlobalName holds the value of the "global_name" field.
	GlobalName string `json:"global_name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"access_token,omitempty"`
	// RefreshToken holds the value of the "refresh_token" field.
	RefreshToken string `json:"refresh_token,omitempty"`
	// Scope holds the value of the "scope" field.
	Scope string `json:"scope,omitempty"`
	// ExpiresIn holds the value of the "expires_in" field.
	ExpiresIn float64 `json:"expires_in,omitempty"`
	// SessionState holds the value of the "session_state" field.
	SessionState string `json:"session_state,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UsersQuery when eager-loading is set.
	Edges        UsersEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UsersEdges holds the relations/edges for other nodes in the graph.
type UsersEdges struct {
	// UsersLinks holds the value of the users_links edge.
	UsersLinks []*Links `json:"users_links,omitempty"`
	// UsersSessions holds the value of the users_sessions edge.
	UsersSessions []*Session `json:"users_sessions,omitempty"`
	// UsersAccounts holds the value of the users_accounts edge.
	UsersAccounts []*Account `json:"users_accounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UsersLinksOrErr returns the UsersLinks value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) UsersLinksOrErr() ([]*Links, error) {
	if e.loadedTypes[0] {
		return e.UsersLinks, nil
	}
	return nil, &NotLoadedError{edge: "users_links"}
}

// UsersSessionsOrErr returns the UsersSessions value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) UsersSessionsOrErr() ([]*Session, error) {
	if e.loadedTypes[1] {
		return e.UsersSessions, nil
	}
	return nil, &NotLoadedError{edge: "users_sessions"}
}

// UsersAccountsOrErr returns the UsersAccounts value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) UsersAccountsOrErr() ([]*Account, error) {
	if e.loadedTypes[2] {
		return e.UsersAccounts, nil
	}
	return nil, &NotLoadedError{edge: "users_accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Users) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case users.FieldDeleted:
			values[i] = new(sql.NullBool)
		case users.FieldExpiresIn:
			values[i] = new(sql.NullFloat64)
		case users.FieldExternalID, users.FieldUsername, users.FieldGlobalName, users.FieldSlug, users.FieldAvatar, users.FieldDescription, users.FieldAccessToken, users.FieldRefreshToken, users.FieldScope, users.FieldSessionState:
			values[i] = new(sql.NullString)
		case users.FieldCreatedAt, users.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case users.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Users fields.
func (u *Users) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case users.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case users.FieldExternalID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_id", values[i])
			} else if value.Valid {
				u.ExternalID = value.String
			}
		case users.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case users.FieldGlobalName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field global_name", values[i])
			} else if value.Valid {
				u.GlobalName = value.String
			}
		case users.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				u.Slug = value.String
			}
		case users.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case users.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				u.Description = value.String
			}
		case users.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				u.AccessToken = value.String
			}
		case users.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				u.RefreshToken = value.String
			}
		case users.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				u.Scope = value.String
			}
		case users.FieldExpiresIn:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field expires_in", values[i])
			} else if value.Valid {
				u.ExpiresIn = value.Float64
			}
		case users.FieldSessionState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field session_state", values[i])
			} else if value.Valid {
				u.SessionState = value.String
			}
		case users.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				u.Deleted = value.Bool
			}
		case users.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case users.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Users.
// This includes values selected through modifiers, order, etc.
func (u *Users) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUsersLinks queries the "users_links" edge of the Users entity.
func (u *Users) QueryUsersLinks() *LinksQuery {
	return NewUsersClient(u.config).QueryUsersLinks(u)
}

// QueryUsersSessions queries the "users_sessions" edge of the Users entity.
func (u *Users) QueryUsersSessions() *SessionQuery {
	return NewUsersClient(u.config).QueryUsersSessions(u)
}

// QueryUsersAccounts queries the "users_accounts" edge of the Users entity.
func (u *Users) QueryUsersAccounts() *AccountQuery {
	return NewUsersClient(u.config).QueryUsersAccounts(u)
}

// Update returns a builder for updating this Users.
// Note that you need to call Users.Unwrap() before calling this method if this Users
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Users) Update() *UsersUpdateOne {
	return NewUsersClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Users entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Users) Unwrap() *Users {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Users is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Users) String() string {
	var builder strings.Builder
	builder.WriteString("Users(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("external_id=")
	builder.WriteString(u.ExternalID)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("global_name=")
	builder.WriteString(u.GlobalName)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(u.Slug)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(u.Description)
	builder.WriteString(", ")
	builder.WriteString("access_token=")
	builder.WriteString(u.AccessToken)
	builder.WriteString(", ")
	builder.WriteString("refresh_token=")
	builder.WriteString(u.RefreshToken)
	builder.WriteString(", ")
	builder.WriteString("scope=")
	builder.WriteString(u.Scope)
	builder.WriteString(", ")
	builder.WriteString("expires_in=")
	builder.WriteString(fmt.Sprintf("%v", u.ExpiresIn))
	builder.WriteString(", ")
	builder.WriteString("session_state=")
	builder.WriteString(u.SessionState)
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", u.Deleted))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UsersSlice is a parsable slice of Users.
type UsersSlice []*Users
