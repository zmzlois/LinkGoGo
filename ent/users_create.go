// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zmzlois/LinkGoGo/ent/links"
	"github.com/zmzlois/LinkGoGo/ent/users"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	mutation *UsersMutation
	hooks    []Hook
}

// SetExternalID sets the "external_id" field.
func (uc *UsersCreate) SetExternalID(s string) *UsersCreate {
	uc.mutation.SetExternalID(s)
	return uc
}

// SetUsername sets the "username" field.
func (uc *UsersCreate) SetUsername(s string) *UsersCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetGlobalName sets the "global_name" field.
func (uc *UsersCreate) SetGlobalName(s string) *UsersCreate {
	uc.mutation.SetGlobalName(s)
	return uc
}

// SetSlug sets the "slug" field.
func (uc *UsersCreate) SetSlug(s string) *UsersCreate {
	uc.mutation.SetSlug(s)
	return uc
}

// SetFirstName sets the "first_name" field.
func (uc *UsersCreate) SetFirstName(s string) *UsersCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetLastName sets the "last_name" field.
func (uc *UsersCreate) SetLastName(s string) *UsersCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UsersCreate) SetEmail(s string) *UsersCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UsersCreate) SetAvatar(s string) *UsersCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetDescription sets the "description" field.
func (uc *UsersCreate) SetDescription(s string) *UsersCreate {
	uc.mutation.SetDescription(s)
	return uc
}

// SetAccessToken sets the "access_token" field.
func (uc *UsersCreate) SetAccessToken(s string) *UsersCreate {
	uc.mutation.SetAccessToken(s)
	return uc
}

// SetRefreshToken sets the "refresh_token" field.
func (uc *UsersCreate) SetRefreshToken(s string) *UsersCreate {
	uc.mutation.SetRefreshToken(s)
	return uc
}

// SetDeleted sets the "deleted" field.
func (uc *UsersCreate) SetDeleted(b bool) *UsersCreate {
	uc.mutation.SetDeleted(b)
	return uc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (uc *UsersCreate) SetNillableDeleted(b *bool) *UsersCreate {
	if b != nil {
		uc.SetDeleted(*b)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UsersCreate) SetCreatedAt(t time.Time) *UsersCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UsersCreate) SetNillableCreatedAt(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UsersCreate) SetUpdatedAt(t time.Time) *UsersCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UsersCreate) SetNillableUpdatedAt(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UsersCreate) SetID(s string) *UsersCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddUsersLinkIDs adds the "users_links" edge to the Links entity by IDs.
func (uc *UsersCreate) AddUsersLinkIDs(ids ...string) *UsersCreate {
	uc.mutation.AddUsersLinkIDs(ids...)
	return uc
}

// AddUsersLinks adds the "users_links" edges to the Links entity.
func (uc *UsersCreate) AddUsersLinks(l ...*Links) *UsersCreate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uc.AddUsersLinkIDs(ids...)
}

// Mutation returns the UsersMutation object of the builder.
func (uc *UsersCreate) Mutation() *UsersMutation {
	return uc.mutation
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UsersCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UsersCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UsersCreate) defaults() {
	if _, ok := uc.mutation.Deleted(); !ok {
		v := users.DefaultDeleted
		uc.mutation.SetDeleted(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := users.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := users.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UsersCreate) check() error {
	if _, ok := uc.mutation.ExternalID(); !ok {
		return &ValidationError{Name: "external_id", err: errors.New(`ent: missing required field "Users.external_id"`)}
	}
	if v, ok := uc.mutation.ExternalID(); ok {
		if err := users.ExternalIDValidator(v); err != nil {
			return &ValidationError{Name: "external_id", err: fmt.Errorf(`ent: validator failed for field "Users.external_id": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Users.username"`)}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := users.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Users.username": %w`, err)}
		}
	}
	if _, ok := uc.mutation.GlobalName(); !ok {
		return &ValidationError{Name: "global_name", err: errors.New(`ent: missing required field "Users.global_name"`)}
	}
	if v, ok := uc.mutation.GlobalName(); ok {
		if err := users.GlobalNameValidator(v); err != nil {
			return &ValidationError{Name: "global_name", err: fmt.Errorf(`ent: validator failed for field "Users.global_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Users.slug"`)}
	}
	if v, ok := uc.mutation.Slug(); ok {
		if err := users.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Users.slug": %w`, err)}
		}
	}
	if _, ok := uc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Users.first_name"`)}
	}
	if v, ok := uc.mutation.FirstName(); ok {
		if err := users.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Users.first_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Users.last_name"`)}
	}
	if v, ok := uc.mutation.LastName(); ok {
		if err := users.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Users.last_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Users.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := users.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Users.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "Users.avatar"`)}
	}
	if v, ok := uc.mutation.Avatar(); ok {
		if err := users.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Users.avatar": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Users.description"`)}
	}
	if v, ok := uc.mutation.Description(); ok {
		if err := users.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Users.description": %w`, err)}
		}
	}
	if _, ok := uc.mutation.AccessToken(); !ok {
		return &ValidationError{Name: "access_token", err: errors.New(`ent: missing required field "Users.access_token"`)}
	}
	if v, ok := uc.mutation.AccessToken(); ok {
		if err := users.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Users.access_token": %w`, err)}
		}
	}
	if _, ok := uc.mutation.RefreshToken(); !ok {
		return &ValidationError{Name: "refresh_token", err: errors.New(`ent: missing required field "Users.refresh_token"`)}
	}
	if v, ok := uc.mutation.RefreshToken(); ok {
		if err := users.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "Users.refresh_token": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Users.deleted"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Users.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Users.updated_at"`)}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := users.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Users.id": %w`, err)}
		}
	}
	return nil
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Users.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UsersCreate) createSpec() (*Users, *sqlgraph.CreateSpec) {
	var (
		_node = &Users{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(users.Table, sqlgraph.NewFieldSpec(users.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.ExternalID(); ok {
		_spec.SetField(users.FieldExternalID, field.TypeString, value)
		_node.ExternalID = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.GlobalName(); ok {
		_spec.SetField(users.FieldGlobalName, field.TypeString, value)
		_node.GlobalName = value
	}
	if value, ok := uc.mutation.Slug(); ok {
		_spec.SetField(users.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.SetField(users.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.SetField(users.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(users.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Description(); ok {
		_spec.SetField(users.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := uc.mutation.AccessToken(); ok {
		_spec.SetField(users.FieldAccessToken, field.TypeString, value)
		_node.AccessToken = value
	}
	if value, ok := uc.mutation.RefreshToken(); ok {
		_spec.SetField(users.FieldRefreshToken, field.TypeString, value)
		_node.RefreshToken = value
	}
	if value, ok := uc.mutation.Deleted(); ok {
		_spec.SetField(users.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(users.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(users.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := uc.mutation.UsersLinksIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UsersCreateBulk is the builder for creating many Users entities in bulk.
type UsersCreateBulk struct {
	config
	err      error
	builders []*UsersCreate
}

// Save creates the Users entities in the database.
func (ucb *UsersCreateBulk) Save(ctx context.Context) ([]*Users, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Users, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UsersCreateBulk) SaveX(ctx context.Context) []*Users {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UsersCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UsersCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
