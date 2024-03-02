// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zmzlois/LinkGoGo/ent/links"
	"github.com/zmzlois/LinkGoGo/ent/users"
)

// LinksCreate is the builder for creating a Links entity.
type LinksCreate struct {
	config
	mutation *LinksMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (lc *LinksCreate) SetUserID(u uuid.UUID) *LinksCreate {
	lc.mutation.SetUserID(u)
	return lc
}

// SetURL sets the "url" field.
func (lc *LinksCreate) SetURL(s string) *LinksCreate {
	lc.mutation.SetURL(s)
	return lc
}

// SetTitle sets the "title" field.
func (lc *LinksCreate) SetTitle(s string) *LinksCreate {
	lc.mutation.SetTitle(s)
	return lc
}

// SetImage sets the "image" field.
func (lc *LinksCreate) SetImage(s string) *LinksCreate {
	lc.mutation.SetImage(s)
	return lc
}

// SetDeleted sets the "deleted" field.
func (lc *LinksCreate) SetDeleted(b bool) *LinksCreate {
	lc.mutation.SetDeleted(b)
	return lc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (lc *LinksCreate) SetNillableDeleted(b *bool) *LinksCreate {
	if b != nil {
		lc.SetDeleted(*b)
	}
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LinksCreate) SetCreatedAt(t time.Time) *LinksCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LinksCreate) SetNillableCreatedAt(t *time.Time) *LinksCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LinksCreate) SetUpdatedAt(t time.Time) *LinksCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LinksCreate) SetNillableUpdatedAt(t *time.Time) *LinksCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LinksCreate) SetID(u uuid.UUID) *LinksCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LinksCreate) SetNillableID(u *uuid.UUID) *LinksCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// SetUser sets the "user" edge to the Users entity.
func (lc *LinksCreate) SetUser(u *Users) *LinksCreate {
	return lc.SetUserID(u.ID)
}

// Mutation returns the LinksMutation object of the builder.
func (lc *LinksCreate) Mutation() *LinksMutation {
	return lc.mutation
}

// Save creates the Links in the database.
func (lc *LinksCreate) Save(ctx context.Context) (*Links, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LinksCreate) SaveX(ctx context.Context) *Links {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LinksCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LinksCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LinksCreate) defaults() {
	if _, ok := lc.mutation.Deleted(); !ok {
		v := links.DefaultDeleted
		lc.mutation.SetDeleted(v)
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := links.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := links.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		v := links.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LinksCreate) check() error {
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Links.user_id"`)}
	}
	if _, ok := lc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Links.url"`)}
	}
	if v, ok := lc.mutation.URL(); ok {
		if err := links.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Links.url": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Links.title"`)}
	}
	if v, ok := lc.mutation.Title(); ok {
		if err := links.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Links.title": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "Links.image"`)}
	}
	if v, ok := lc.mutation.Image(); ok {
		if err := links.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Links.image": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Links.deleted"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Links.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Links.updated_at"`)}
	}
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Links.user"`)}
	}
	return nil
}

func (lc *LinksCreate) sqlSave(ctx context.Context) (*Links, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LinksCreate) createSpec() (*Links, *sqlgraph.CreateSpec) {
	var (
		_node = &Links{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(links.Table, sqlgraph.NewFieldSpec(links.FieldID, field.TypeUUID))
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.URL(); ok {
		_spec.SetField(links.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := lc.mutation.Title(); ok {
		_spec.SetField(links.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := lc.mutation.Image(); ok {
		_spec.SetField(links.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := lc.mutation.Deleted(); ok {
		_spec.SetField(links.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(links.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(links.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := lc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   links.UserTable,
			Columns: []string{links.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LinksCreateBulk is the builder for creating many Links entities in bulk.
type LinksCreateBulk struct {
	config
	err      error
	builders []*LinksCreate
}

// Save creates the Links entities in the database.
func (lcb *LinksCreateBulk) Save(ctx context.Context) ([]*Links, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Links, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinksMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LinksCreateBulk) SaveX(ctx context.Context) []*Links {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LinksCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LinksCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
