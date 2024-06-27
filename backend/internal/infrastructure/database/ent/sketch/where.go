// Code generated by ent, DO NOT EDIT.

package sketch

import (
	"myapp/internal/infrastructure/database/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Sketch {
	return predicate.Sketch(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Sketch {
	return predicate.Sketch(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Sketch {
	return predicate.Sketch(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Sketch {
	return predicate.Sketch(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Sketch {
	return predicate.Sketch(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Sketch {
	return predicate.Sketch(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Sketch {
	return predicate.Sketch(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldUserID, v))
}

// ImageName applies equality check predicate on the "image_name" field. It's identical to ImageNameEQ.
func ImageName(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldImageName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Sketch {
	return predicate.Sketch(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Sketch {
	return predicate.Sketch(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Sketch {
	return predicate.Sketch(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Sketch {
	return predicate.Sketch(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.Sketch {
	return predicate.Sketch(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.Sketch {
	return predicate.Sketch(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.Sketch {
	return predicate.Sketch(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.Sketch {
	return predicate.Sketch(sql.FieldLTE(FieldUserID, v))
}

// ImageNameEQ applies the EQ predicate on the "image_name" field.
func ImageNameEQ(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldEQ(FieldImageName, v))
}

// ImageNameNEQ applies the NEQ predicate on the "image_name" field.
func ImageNameNEQ(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldNEQ(FieldImageName, v))
}

// ImageNameIn applies the In predicate on the "image_name" field.
func ImageNameIn(vs ...string) predicate.Sketch {
	return predicate.Sketch(sql.FieldIn(FieldImageName, vs...))
}

// ImageNameNotIn applies the NotIn predicate on the "image_name" field.
func ImageNameNotIn(vs ...string) predicate.Sketch {
	return predicate.Sketch(sql.FieldNotIn(FieldImageName, vs...))
}

// ImageNameGT applies the GT predicate on the "image_name" field.
func ImageNameGT(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldGT(FieldImageName, v))
}

// ImageNameGTE applies the GTE predicate on the "image_name" field.
func ImageNameGTE(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldGTE(FieldImageName, v))
}

// ImageNameLT applies the LT predicate on the "image_name" field.
func ImageNameLT(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldLT(FieldImageName, v))
}

// ImageNameLTE applies the LTE predicate on the "image_name" field.
func ImageNameLTE(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldLTE(FieldImageName, v))
}

// ImageNameContains applies the Contains predicate on the "image_name" field.
func ImageNameContains(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldContains(FieldImageName, v))
}

// ImageNameHasPrefix applies the HasPrefix predicate on the "image_name" field.
func ImageNameHasPrefix(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldHasPrefix(FieldImageName, v))
}

// ImageNameHasSuffix applies the HasSuffix predicate on the "image_name" field.
func ImageNameHasSuffix(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldHasSuffix(FieldImageName, v))
}

// ImageNameEqualFold applies the EqualFold predicate on the "image_name" field.
func ImageNameEqualFold(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldEqualFold(FieldImageName, v))
}

// ImageNameContainsFold applies the ContainsFold predicate on the "image_name" field.
func ImageNameContainsFold(v string) predicate.Sketch {
	return predicate.Sketch(sql.FieldContainsFold(FieldImageName, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Sketch) predicate.Sketch {
	return predicate.Sketch(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Sketch) predicate.Sketch {
	return predicate.Sketch(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Sketch) predicate.Sketch {
	return predicate.Sketch(sql.NotPredicates(p))
}
