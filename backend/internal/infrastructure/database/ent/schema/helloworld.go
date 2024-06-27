package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HelloWorld holds the schema definition for the HelloWorld entity.
type HelloWorld struct {
	ent.Schema
}

// Mixin of the HelloWorld.
func (HelloWorld) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the HelloWorld.
func (HelloWorld) Fields() []ent.Field {
	return []ent.Field{
		field.String("lang"),
		field.String("message"),
	}
}

// Edges of the HelloWorld.
func (HelloWorld) Edges() []ent.Edge {
	return nil
}
