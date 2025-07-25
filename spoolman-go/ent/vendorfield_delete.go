// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"spoolman-go/ent/predicate"
	"spoolman-go/ent/vendorfield"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VendorFieldDelete is the builder for deleting a VendorField entity.
type VendorFieldDelete struct {
	config
	hooks    []Hook
	mutation *VendorFieldMutation
}

// Where appends a list predicates to the VendorFieldDelete builder.
func (vfd *VendorFieldDelete) Where(ps ...predicate.VendorField) *VendorFieldDelete {
	vfd.mutation.Where(ps...)
	return vfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vfd *VendorFieldDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vfd.sqlExec, vfd.mutation, vfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vfd *VendorFieldDelete) ExecX(ctx context.Context) int {
	n, err := vfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vfd *VendorFieldDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(vendorfield.Table, sqlgraph.NewFieldSpec(vendorfield.FieldID, field.TypeInt))
	if ps := vfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vfd.mutation.done = true
	return affected, err
}

// VendorFieldDeleteOne is the builder for deleting a single VendorField entity.
type VendorFieldDeleteOne struct {
	vfd *VendorFieldDelete
}

// Where appends a list predicates to the VendorFieldDelete builder.
func (vfdo *VendorFieldDeleteOne) Where(ps ...predicate.VendorField) *VendorFieldDeleteOne {
	vfdo.vfd.mutation.Where(ps...)
	return vfdo
}

// Exec executes the deletion query.
func (vfdo *VendorFieldDeleteOne) Exec(ctx context.Context) error {
	n, err := vfdo.vfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vendorfield.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vfdo *VendorFieldDeleteOne) ExecX(ctx context.Context) {
	if err := vfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
