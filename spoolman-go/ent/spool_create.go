// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"spoolman-go/ent/filament"
	"spoolman-go/ent/spool"
	"spoolman-go/ent/spoolfield"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpoolCreate is the builder for creating a Spool entity.
type SpoolCreate struct {
	config
	mutation *SpoolMutation
	hooks    []Hook
}

// SetRegistered sets the "registered" field.
func (sc *SpoolCreate) SetRegistered(t time.Time) *SpoolCreate {
	sc.mutation.SetRegistered(t)
	return sc
}

// SetFirstUsed sets the "first_used" field.
func (sc *SpoolCreate) SetFirstUsed(t time.Time) *SpoolCreate {
	sc.mutation.SetFirstUsed(t)
	return sc
}

// SetNillableFirstUsed sets the "first_used" field if the given value is not nil.
func (sc *SpoolCreate) SetNillableFirstUsed(t *time.Time) *SpoolCreate {
	if t != nil {
		sc.SetFirstUsed(*t)
	}
	return sc
}

// SetLastUsed sets the "last_used" field.
func (sc *SpoolCreate) SetLastUsed(t time.Time) *SpoolCreate {
	sc.mutation.SetLastUsed(t)
	return sc
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (sc *SpoolCreate) SetNillableLastUsed(t *time.Time) *SpoolCreate {
	if t != nil {
		sc.SetLastUsed(*t)
	}
	return sc
}

// SetPrice sets the "price" field.
func (sc *SpoolCreate) SetPrice(f float64) *SpoolCreate {
	sc.mutation.SetPrice(f)
	return sc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (sc *SpoolCreate) SetNillablePrice(f *float64) *SpoolCreate {
	if f != nil {
		sc.SetPrice(*f)
	}
	return sc
}

// SetFilamentID sets the "filament_id" field.
func (sc *SpoolCreate) SetFilamentID(i int) *SpoolCreate {
	sc.mutation.SetFilamentID(i)
	return sc
}

// SetInitialWeight sets the "initial_weight" field.
func (sc *SpoolCreate) SetInitialWeight(f float64) *SpoolCreate {
	sc.mutation.SetInitialWeight(f)
	return sc
}

// SetNillableInitialWeight sets the "initial_weight" field if the given value is not nil.
func (sc *SpoolCreate) SetNillableInitialWeight(f *float64) *SpoolCreate {
	if f != nil {
		sc.SetInitialWeight(*f)
	}
	return sc
}

// SetSpoolWeight sets the "spool_weight" field.
func (sc *SpoolCreate) SetSpoolWeight(f float64) *SpoolCreate {
	sc.mutation.SetSpoolWeight(f)
	return sc
}

// SetNillableSpoolWeight sets the "spool_weight" field if the given value is not nil.
func (sc *SpoolCreate) SetNillableSpoolWeight(f *float64) *SpoolCreate {
	if f != nil {
		sc.SetSpoolWeight(*f)
	}
	return sc
}

// SetUsedWeight sets the "used_weight" field.
func (sc *SpoolCreate) SetUsedWeight(f float64) *SpoolCreate {
	sc.mutation.SetUsedWeight(f)
	return sc
}

// SetLocation sets the "location" field.
func (sc *SpoolCreate) SetLocation(s string) *SpoolCreate {
	sc.mutation.SetLocation(s)
	return sc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (sc *SpoolCreate) SetNillableLocation(s *string) *SpoolCreate {
	if s != nil {
		sc.SetLocation(*s)
	}
	return sc
}

// SetLotNr sets the "lot_nr" field.
func (sc *SpoolCreate) SetLotNr(s string) *SpoolCreate {
	sc.mutation.SetLotNr(s)
	return sc
}

// SetNillableLotNr sets the "lot_nr" field if the given value is not nil.
func (sc *SpoolCreate) SetNillableLotNr(s *string) *SpoolCreate {
	if s != nil {
		sc.SetLotNr(*s)
	}
	return sc
}

// SetComment sets the "comment" field.
func (sc *SpoolCreate) SetComment(s string) *SpoolCreate {
	sc.mutation.SetComment(s)
	return sc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (sc *SpoolCreate) SetNillableComment(s *string) *SpoolCreate {
	if s != nil {
		sc.SetComment(*s)
	}
	return sc
}

// SetArchived sets the "archived" field.
func (sc *SpoolCreate) SetArchived(b bool) *SpoolCreate {
	sc.mutation.SetArchived(b)
	return sc
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (sc *SpoolCreate) SetNillableArchived(b *bool) *SpoolCreate {
	if b != nil {
		sc.SetArchived(*b)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SpoolCreate) SetID(i int) *SpoolCreate {
	sc.mutation.SetID(i)
	return sc
}

// SetFilament sets the "filament" edge to the Filament entity.
func (sc *SpoolCreate) SetFilament(f *Filament) *SpoolCreate {
	return sc.SetFilamentID(f.ID)
}

// AddExtraIDs adds the "extra" edge to the SpoolField entity by IDs.
func (sc *SpoolCreate) AddExtraIDs(ids ...int) *SpoolCreate {
	sc.mutation.AddExtraIDs(ids...)
	return sc
}

// AddExtra adds the "extra" edges to the SpoolField entity.
func (sc *SpoolCreate) AddExtra(s ...*SpoolField) *SpoolCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddExtraIDs(ids...)
}

// Mutation returns the SpoolMutation object of the builder.
func (sc *SpoolCreate) Mutation() *SpoolMutation {
	return sc.mutation
}

// Save creates the Spool in the database.
func (sc *SpoolCreate) Save(ctx context.Context) (*Spool, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SpoolCreate) SaveX(ctx context.Context) *Spool {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SpoolCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SpoolCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SpoolCreate) check() error {
	if _, ok := sc.mutation.Registered(); !ok {
		return &ValidationError{Name: "registered", err: errors.New(`ent: missing required field "Spool.registered"`)}
	}
	if _, ok := sc.mutation.FilamentID(); !ok {
		return &ValidationError{Name: "filament_id", err: errors.New(`ent: missing required field "Spool.filament_id"`)}
	}
	if _, ok := sc.mutation.UsedWeight(); !ok {
		return &ValidationError{Name: "used_weight", err: errors.New(`ent: missing required field "Spool.used_weight"`)}
	}
	if v, ok := sc.mutation.Location(); ok {
		if err := spool.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "Spool.location": %w`, err)}
		}
	}
	if v, ok := sc.mutation.LotNr(); ok {
		if err := spool.LotNrValidator(v); err != nil {
			return &ValidationError{Name: "lot_nr", err: fmt.Errorf(`ent: validator failed for field "Spool.lot_nr": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Comment(); ok {
		if err := spool.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`ent: validator failed for field "Spool.comment": %w`, err)}
		}
	}
	if v, ok := sc.mutation.ID(); ok {
		if err := spool.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Spool.id": %w`, err)}
		}
	}
	if len(sc.mutation.FilamentIDs()) == 0 {
		return &ValidationError{Name: "filament", err: errors.New(`ent: missing required edge "Spool.filament"`)}
	}
	return nil
}

func (sc *SpoolCreate) sqlSave(ctx context.Context) (*Spool, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SpoolCreate) createSpec() (*Spool, *sqlgraph.CreateSpec) {
	var (
		_node = &Spool{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(spool.Table, sqlgraph.NewFieldSpec(spool.FieldID, field.TypeInt))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Registered(); ok {
		_spec.SetField(spool.FieldRegistered, field.TypeTime, value)
		_node.Registered = value
	}
	if value, ok := sc.mutation.FirstUsed(); ok {
		_spec.SetField(spool.FieldFirstUsed, field.TypeTime, value)
		_node.FirstUsed = value
	}
	if value, ok := sc.mutation.LastUsed(); ok {
		_spec.SetField(spool.FieldLastUsed, field.TypeTime, value)
		_node.LastUsed = value
	}
	if value, ok := sc.mutation.Price(); ok {
		_spec.SetField(spool.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := sc.mutation.InitialWeight(); ok {
		_spec.SetField(spool.FieldInitialWeight, field.TypeFloat64, value)
		_node.InitialWeight = value
	}
	if value, ok := sc.mutation.SpoolWeight(); ok {
		_spec.SetField(spool.FieldSpoolWeight, field.TypeFloat64, value)
		_node.SpoolWeight = value
	}
	if value, ok := sc.mutation.UsedWeight(); ok {
		_spec.SetField(spool.FieldUsedWeight, field.TypeFloat64, value)
		_node.UsedWeight = value
	}
	if value, ok := sc.mutation.Location(); ok {
		_spec.SetField(spool.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := sc.mutation.LotNr(); ok {
		_spec.SetField(spool.FieldLotNr, field.TypeString, value)
		_node.LotNr = value
	}
	if value, ok := sc.mutation.Comment(); ok {
		_spec.SetField(spool.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if value, ok := sc.mutation.Archived(); ok {
		_spec.SetField(spool.FieldArchived, field.TypeBool, value)
		_node.Archived = value
	}
	if nodes := sc.mutation.FilamentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   spool.FilamentTable,
			Columns: []string{spool.FilamentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(filament.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FilamentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ExtraIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   spool.ExtraTable,
			Columns: []string{spool.ExtraColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spoolfield.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SpoolCreateBulk is the builder for creating many Spool entities in bulk.
type SpoolCreateBulk struct {
	config
	err      error
	builders []*SpoolCreate
}

// Save creates the Spool entities in the database.
func (scb *SpoolCreateBulk) Save(ctx context.Context) ([]*Spool, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Spool, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SpoolMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SpoolCreateBulk) SaveX(ctx context.Context) []*Spool {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SpoolCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SpoolCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
