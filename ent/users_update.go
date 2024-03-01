// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zmzlois/LinkGoGo/ent/links"
	"github.com/zmzlois/LinkGoGo/ent/predicate"
	"github.com/zmzlois/LinkGoGo/ent/users"
)

// UsersUpdate is the builder for updating Users entities.
type UsersUpdate struct {
	config
	hooks    []Hook
	mutation *UsersMutation
}

// Where appends a list predicates to the UsersUpdate builder.
func (uu *UsersUpdate) Where(ps ...predicate.Users) *UsersUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetExternalID sets the "external_id" field.
func (uu *UsersUpdate) SetExternalID(s string) *UsersUpdate {
	uu.mutation.SetExternalID(s)
	return uu
}

// SetNillableExternalID sets the "external_id" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableExternalID(s *string) *UsersUpdate {
	if s != nil {
		uu.SetExternalID(*s)
	}
	return uu
}

// SetUsername sets the "username" field.
func (uu *UsersUpdate) SetUsername(s string) *UsersUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableUsername(s *string) *UsersUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetGlobalName sets the "global_name" field.
func (uu *UsersUpdate) SetGlobalName(s string) *UsersUpdate {
	uu.mutation.SetGlobalName(s)
	return uu
}

// SetNillableGlobalName sets the "global_name" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableGlobalName(s *string) *UsersUpdate {
	if s != nil {
		uu.SetGlobalName(*s)
	}
	return uu
}

// SetSlug sets the "slug" field.
func (uu *UsersUpdate) SetSlug(s string) *UsersUpdate {
	uu.mutation.SetSlug(s)
	return uu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableSlug(s *string) *UsersUpdate {
	if s != nil {
		uu.SetSlug(*s)
	}
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UsersUpdate) SetFirstName(s string) *UsersUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableFirstName(s *string) *UsersUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UsersUpdate) SetLastName(s string) *UsersUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableLastName(s *string) *UsersUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UsersUpdate) SetEmail(s string) *UsersUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableEmail(s *string) *UsersUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UsersUpdate) SetAvatar(s string) *UsersUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableAvatar(s *string) *UsersUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// SetDescription sets the "description" field.
func (uu *UsersUpdate) SetDescription(s string) *UsersUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableDescription(s *string) *UsersUpdate {
	if s != nil {
		uu.SetDescription(*s)
	}
	return uu
}

// SetAccessToken sets the "access_token" field.
func (uu *UsersUpdate) SetAccessToken(s string) *UsersUpdate {
	uu.mutation.SetAccessToken(s)
	return uu
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableAccessToken(s *string) *UsersUpdate {
	if s != nil {
		uu.SetAccessToken(*s)
	}
	return uu
}

// SetRefreshToken sets the "refresh_token" field.
func (uu *UsersUpdate) SetRefreshToken(s string) *UsersUpdate {
	uu.mutation.SetRefreshToken(s)
	return uu
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableRefreshToken(s *string) *UsersUpdate {
	if s != nil {
		uu.SetRefreshToken(*s)
	}
	return uu
}

// SetDeleted sets the "deleted" field.
func (uu *UsersUpdate) SetDeleted(b bool) *UsersUpdate {
	uu.mutation.SetDeleted(b)
	return uu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableDeleted(b *bool) *UsersUpdate {
	if b != nil {
		uu.SetDeleted(*b)
	}
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UsersUpdate) SetCreatedAt(t time.Time) *UsersUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableCreatedAt(t *time.Time) *UsersUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UsersUpdate) SetUpdatedAt(t time.Time) *UsersUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// AddUsersLinkIDs adds the "users_links" edge to the Links entity by IDs.
func (uu *UsersUpdate) AddUsersLinkIDs(ids ...string) *UsersUpdate {
	uu.mutation.AddUsersLinkIDs(ids...)
	return uu
}

// AddUsersLinks adds the "users_links" edges to the Links entity.
func (uu *UsersUpdate) AddUsersLinks(l ...*Links) *UsersUpdate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.AddUsersLinkIDs(ids...)
}

// Mutation returns the UsersMutation object of the builder.
func (uu *UsersUpdate) Mutation() *UsersMutation {
	return uu.mutation
}

// ClearUsersLinks clears all "users_links" edges to the Links entity.
func (uu *UsersUpdate) ClearUsersLinks() *UsersUpdate {
	uu.mutation.ClearUsersLinks()
	return uu
}

// RemoveUsersLinkIDs removes the "users_links" edge to Links entities by IDs.
func (uu *UsersUpdate) RemoveUsersLinkIDs(ids ...string) *UsersUpdate {
	uu.mutation.RemoveUsersLinkIDs(ids...)
	return uu
}

// RemoveUsersLinks removes "users_links" edges to Links entities.
func (uu *UsersUpdate) RemoveUsersLinks(l ...*Links) *UsersUpdate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.RemoveUsersLinkIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UsersUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UsersUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UsersUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UsersUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UsersUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := users.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UsersUpdate) check() error {
	if v, ok := uu.mutation.ExternalID(); ok {
		if err := users.ExternalIDValidator(v); err != nil {
			return &ValidationError{Name: "external_id", err: fmt.Errorf(`ent: validator failed for field "Users.external_id": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Username(); ok {
		if err := users.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Users.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.GlobalName(); ok {
		if err := users.GlobalNameValidator(v); err != nil {
			return &ValidationError{Name: "global_name", err: fmt.Errorf(`ent: validator failed for field "Users.global_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Slug(); ok {
		if err := users.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Users.slug": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FirstName(); ok {
		if err := users.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Users.first_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := users.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Users.last_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := users.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Users.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Avatar(); ok {
		if err := users.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Users.avatar": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Description(); ok {
		if err := users.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Users.description": %w`, err)}
		}
	}
	if v, ok := uu.mutation.AccessToken(); ok {
		if err := users.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Users.access_token": %w`, err)}
		}
	}
	if v, ok := uu.mutation.RefreshToken(); ok {
		if err := users.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "Users.refresh_token": %w`, err)}
		}
	}
	return nil
}

func (uu *UsersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeString))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.ExternalID(); ok {
		_spec.SetField(users.FieldExternalID, field.TypeString, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.GlobalName(); ok {
		_spec.SetField(users.FieldGlobalName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Slug(); ok {
		_spec.SetField(users.FieldSlug, field.TypeString, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(users.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(users.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(users.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.SetField(users.FieldDescription, field.TypeString, value)
	}
	if value, ok := uu.mutation.AccessToken(); ok {
		_spec.SetField(users.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := uu.mutation.RefreshToken(); ok {
		_spec.SetField(users.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := uu.mutation.Deleted(); ok {
		_spec.SetField(users.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(users.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(users.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.UsersLinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.UsersLinksTable,
			Columns: []string{users.UsersLinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(links.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUsersLinksIDs(); len(nodes) > 0 && !uu.mutation.UsersLinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.UsersLinksTable,
			Columns: []string{users.UsersLinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(links.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UsersLinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.UsersLinksTable,
			Columns: []string{users.UsersLinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(links.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UsersUpdateOne is the builder for updating a single Users entity.
type UsersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersMutation
}

// SetExternalID sets the "external_id" field.
func (uuo *UsersUpdateOne) SetExternalID(s string) *UsersUpdateOne {
	uuo.mutation.SetExternalID(s)
	return uuo
}

// SetNillableExternalID sets the "external_id" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableExternalID(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetExternalID(*s)
	}
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UsersUpdateOne) SetUsername(s string) *UsersUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableUsername(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetGlobalName sets the "global_name" field.
func (uuo *UsersUpdateOne) SetGlobalName(s string) *UsersUpdateOne {
	uuo.mutation.SetGlobalName(s)
	return uuo
}

// SetNillableGlobalName sets the "global_name" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableGlobalName(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetGlobalName(*s)
	}
	return uuo
}

// SetSlug sets the "slug" field.
func (uuo *UsersUpdateOne) SetSlug(s string) *UsersUpdateOne {
	uuo.mutation.SetSlug(s)
	return uuo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableSlug(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetSlug(*s)
	}
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UsersUpdateOne) SetFirstName(s string) *UsersUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableFirstName(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UsersUpdateOne) SetLastName(s string) *UsersUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableLastName(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UsersUpdateOne) SetEmail(s string) *UsersUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableEmail(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UsersUpdateOne) SetAvatar(s string) *UsersUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableAvatar(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UsersUpdateOne) SetDescription(s string) *UsersUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableDescription(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetDescription(*s)
	}
	return uuo
}

// SetAccessToken sets the "access_token" field.
func (uuo *UsersUpdateOne) SetAccessToken(s string) *UsersUpdateOne {
	uuo.mutation.SetAccessToken(s)
	return uuo
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableAccessToken(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetAccessToken(*s)
	}
	return uuo
}

// SetRefreshToken sets the "refresh_token" field.
func (uuo *UsersUpdateOne) SetRefreshToken(s string) *UsersUpdateOne {
	uuo.mutation.SetRefreshToken(s)
	return uuo
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableRefreshToken(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetRefreshToken(*s)
	}
	return uuo
}

// SetDeleted sets the "deleted" field.
func (uuo *UsersUpdateOne) SetDeleted(b bool) *UsersUpdateOne {
	uuo.mutation.SetDeleted(b)
	return uuo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableDeleted(b *bool) *UsersUpdateOne {
	if b != nil {
		uuo.SetDeleted(*b)
	}
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UsersUpdateOne) SetCreatedAt(t time.Time) *UsersUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableCreatedAt(t *time.Time) *UsersUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UsersUpdateOne) SetUpdatedAt(t time.Time) *UsersUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// AddUsersLinkIDs adds the "users_links" edge to the Links entity by IDs.
func (uuo *UsersUpdateOne) AddUsersLinkIDs(ids ...string) *UsersUpdateOne {
	uuo.mutation.AddUsersLinkIDs(ids...)
	return uuo
}

// AddUsersLinks adds the "users_links" edges to the Links entity.
func (uuo *UsersUpdateOne) AddUsersLinks(l ...*Links) *UsersUpdateOne {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.AddUsersLinkIDs(ids...)
}

// Mutation returns the UsersMutation object of the builder.
func (uuo *UsersUpdateOne) Mutation() *UsersMutation {
	return uuo.mutation
}

// ClearUsersLinks clears all "users_links" edges to the Links entity.
func (uuo *UsersUpdateOne) ClearUsersLinks() *UsersUpdateOne {
	uuo.mutation.ClearUsersLinks()
	return uuo
}

// RemoveUsersLinkIDs removes the "users_links" edge to Links entities by IDs.
func (uuo *UsersUpdateOne) RemoveUsersLinkIDs(ids ...string) *UsersUpdateOne {
	uuo.mutation.RemoveUsersLinkIDs(ids...)
	return uuo
}

// RemoveUsersLinks removes "users_links" edges to Links entities.
func (uuo *UsersUpdateOne) RemoveUsersLinks(l ...*Links) *UsersUpdateOne {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.RemoveUsersLinkIDs(ids...)
}

// Where appends a list predicates to the UsersUpdate builder.
func (uuo *UsersUpdateOne) Where(ps ...predicate.Users) *UsersUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UsersUpdateOne) Select(field string, fields ...string) *UsersUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Users entity.
func (uuo *UsersUpdateOne) Save(ctx context.Context) (*Users, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UsersUpdateOne) SaveX(ctx context.Context) *Users {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UsersUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UsersUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UsersUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := users.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UsersUpdateOne) check() error {
	if v, ok := uuo.mutation.ExternalID(); ok {
		if err := users.ExternalIDValidator(v); err != nil {
			return &ValidationError{Name: "external_id", err: fmt.Errorf(`ent: validator failed for field "Users.external_id": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Username(); ok {
		if err := users.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Users.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.GlobalName(); ok {
		if err := users.GlobalNameValidator(v); err != nil {
			return &ValidationError{Name: "global_name", err: fmt.Errorf(`ent: validator failed for field "Users.global_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Slug(); ok {
		if err := users.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Users.slug": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := users.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Users.first_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := users.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Users.last_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := users.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Users.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Avatar(); ok {
		if err := users.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Users.avatar": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Description(); ok {
		if err := users.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Users.description": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.AccessToken(); ok {
		if err := users.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Users.access_token": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.RefreshToken(); ok {
		if err := users.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "Users.refresh_token": %w`, err)}
		}
	}
	return nil
}

func (uuo *UsersUpdateOne) sqlSave(ctx context.Context) (_node *Users, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeString))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Users.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for _, f := range fields {
			if !users.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.ExternalID(); ok {
		_spec.SetField(users.FieldExternalID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.GlobalName(); ok {
		_spec.SetField(users.FieldGlobalName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Slug(); ok {
		_spec.SetField(users.FieldSlug, field.TypeString, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(users.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(users.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(users.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.SetField(users.FieldDescription, field.TypeString, value)
	}
	if value, ok := uuo.mutation.AccessToken(); ok {
		_spec.SetField(users.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := uuo.mutation.RefreshToken(); ok {
		_spec.SetField(users.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Deleted(); ok {
		_spec.SetField(users.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(users.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(users.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.UsersLinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.UsersLinksTable,
			Columns: []string{users.UsersLinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(links.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUsersLinksIDs(); len(nodes) > 0 && !uuo.mutation.UsersLinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.UsersLinksTable,
			Columns: []string{users.UsersLinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(links.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UsersLinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.UsersLinksTable,
			Columns: []string{users.UsersLinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(links.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Users{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
