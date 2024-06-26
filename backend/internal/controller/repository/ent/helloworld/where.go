// Code generated by ent, DO NOT EDIT.

package helloworld

import (
	"myapp/internal/controller/repository/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldUpdatedAt, v))
}

// Lang applies equality check predicate on the "lang" field. It's identical to LangEQ.
func Lang(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldLang, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldMessage, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLTE(FieldUpdatedAt, v))
}

// LangEQ applies the EQ predicate on the "lang" field.
func LangEQ(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldLang, v))
}

// LangNEQ applies the NEQ predicate on the "lang" field.
func LangNEQ(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNEQ(FieldLang, v))
}

// LangIn applies the In predicate on the "lang" field.
func LangIn(vs ...string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldIn(FieldLang, vs...))
}

// LangNotIn applies the NotIn predicate on the "lang" field.
func LangNotIn(vs ...string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNotIn(FieldLang, vs...))
}

// LangGT applies the GT predicate on the "lang" field.
func LangGT(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGT(FieldLang, v))
}

// LangGTE applies the GTE predicate on the "lang" field.
func LangGTE(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGTE(FieldLang, v))
}

// LangLT applies the LT predicate on the "lang" field.
func LangLT(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLT(FieldLang, v))
}

// LangLTE applies the LTE predicate on the "lang" field.
func LangLTE(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLTE(FieldLang, v))
}

// LangContains applies the Contains predicate on the "lang" field.
func LangContains(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldContains(FieldLang, v))
}

// LangHasPrefix applies the HasPrefix predicate on the "lang" field.
func LangHasPrefix(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldHasPrefix(FieldLang, v))
}

// LangHasSuffix applies the HasSuffix predicate on the "lang" field.
func LangHasSuffix(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldHasSuffix(FieldLang, v))
}

// LangEqualFold applies the EqualFold predicate on the "lang" field.
func LangEqualFold(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEqualFold(FieldLang, v))
}

// LangContainsFold applies the ContainsFold predicate on the "lang" field.
func LangContainsFold(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldContainsFold(FieldLang, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.HelloWorld {
	return predicate.HelloWorld(sql.FieldContainsFold(FieldMessage, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HelloWorld) predicate.HelloWorld {
	return predicate.HelloWorld(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HelloWorld) predicate.HelloWorld {
	return predicate.HelloWorld(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.HelloWorld) predicate.HelloWorld {
	return predicate.HelloWorld(sql.NotPredicates(p))
}
