// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"myapp/internal/controller/repository/ent/predicate"
	"myapp/internal/controller/repository/ent/sketch"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SketchDelete is the builder for deleting a Sketch entity.
type SketchDelete struct {
	config
	hooks    []Hook
	mutation *SketchMutation
}

// Where appends a list predicates to the SketchDelete builder.
func (sd *SketchDelete) Where(ps ...predicate.Sketch) *SketchDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SketchDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SketchDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SketchDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sketch.Table, sqlgraph.NewFieldSpec(sketch.FieldID, field.TypeInt))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SketchDeleteOne is the builder for deleting a single Sketch entity.
type SketchDeleteOne struct {
	sd *SketchDelete
}

// Where appends a list predicates to the SketchDelete builder.
func (sdo *SketchDeleteOne) Where(ps ...predicate.Sketch) *SketchDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SketchDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sketch.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SketchDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
